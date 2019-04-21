package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"sync"
	"text/template"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

type context struct {
	parent *context
	msg    *messageDef
}

func (ctx *context) lookup(name string) *messageDef {
	camelName := strcase.ToCamel(name)
	for _, msg := range ctx.msg.Messages {
		if camelName == msg.Name {
			return msg
		}
	}

	if ctx.parent != nil {
		return ctx.parent.lookup(name)
	}

	return nil
}

type fieldType interface {
	protoType(ctx *context) (string, string, error)
}

type valueType struct {
	Type        schema.ValueType
	ElementType fieldType
}

var setErrorOnce sync.Once

func (ty *valueType) protoType(ctx *context) (ptype string, validation string, err error) {

	switch ty.Type {
	case schema.TypeBool:
		ptype = "bool"
	case schema.TypeInt:
		ptype = "int64"
	case schema.TypeFloat:
		ptype = "double"
	case schema.TypeString:
		ptype = "string"
	case schema.TypeSet:
		validation = "[(validate.rules).repeated.unique = true]"
		if _, ok := ty.ElementType.(*valueType); !ok {
			setErrorOnce.Do(func() {
				log.Infof("TODO: gogoprotobuf unique constraint only works for scalar types, not %#v", ty.ElementType)
			})
			validation = ""
		}
		fallthrough
	case schema.TypeList:
		ptype = "repeated "
		// TODO fix this after fixing resources below
		if ty.ElementType != nil {
			eType, suffix, err := ty.ElementType.protoType(ctx)
			if err != nil {
				return "", "", errors.Wrapf(err, "%#v.protoType", ty.ElementType)
			}
			if suffix != "" {
				return "", "", fmt.Errorf("%#v.protoType: didn't expect Elem to require suffix", ty.ElementType)
			}
			ptype += eType
		} else {
			log.Warnf("nil repeated type for %#v", ty)
		}
	case schema.TypeMap:
		var valueType, suffix string
		if ty.ElementType != nil {
			valueType, suffix, err = ty.ElementType.protoType(ctx)
			if err != nil {
				return "", "", errors.Wrapf(err, "%#v.protoType", ty.ElementType)
			}
			if suffix != "" {
				return "", "", fmt.Errorf("%#v.protoType: didn't expect Elem to require suffix", ty.ElementType)
			}
		} else {
			log.Warnf("nil map type for %#v", ty)
		}
		ptype = fmt.Sprintf("map<string, %s>", valueType)
	default:
		return "", "", errors.Errorf(":ohno: unknown Schema.Type %#v", ty.Type)
	}

	return ptype, validation, nil
}

type referenceType struct {
	Name string
}

func (ty *referenceType) protoType(ctx *context) (string, string, error) {
	msg := ctx.lookup(ty.Name)
	if msg == nil {
		return "", "", fmt.Errorf("reference to type not in context: %s", ty.Name)
	}
	return ty.Name, "", nil
}

// keys returns a sorted list of keys for a given schema map
func keys(m map[string]*schema.Schema) []string {
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}

type fieldDef struct {
	Name     string
	Number   uint
	Type     fieldType
	Optional bool
	MinItems *uint
	MaxItems *uint
}

func (f *fieldDef) protoField(ctx *context) (string, error) {
	// TODO: required annotation

	ty, suffix, err := f.Type.protoType(ctx)
	if err != nil {
		return "", err
	}
	if !strings.HasPrefix(suffix, " ") {
		suffix = " " + suffix
	}
	if !strings.HasSuffix(suffix, ";") {
		suffix += ";"
	}
	name := strcase.ToSnake(f.Name)
	return fmt.Sprintf("%s %s = %d%s", ty, name, f.Number, suffix), nil
}

type messageDef struct {
	Name string
	// messages can have sub-messages, referenced by fields
	Messages []*messageDef
	Fields   []*fieldDef
}

func newMessage(name string, resource *schema.Resource) (*messageDef, error) {
	msg := &messageDef{
		Name:   strcase.ToCamel(name),
		Fields: make([]*fieldDef, 0, len(resource.Schema)),
	}

	fieldNames := keys(resource.Schema)
	for i, name := range fieldNames {
		s := resource.Schema[name]

		def := &fieldDef{
			Name:     name,
			Number:   uint(i) + 1,
			Optional: s.Optional,
		}

		var elementType fieldType

		// first check for nil interface value (different from
		// a typed-nil in the interface)
		if elem := s.Elem; elem != nil {
			switch e := elem.(type) {
			case *schema.Resource:
				if e == nil {
					break
				}
				submsg, err := newMessage(name, e)
				if err != nil {
					return nil, errors.Wrapf(err, "newMessage(%s)", name)
				}
				msg.Messages = append(msg.Messages, submsg)
				elementType = &referenceType{submsg.Name}
			case *schema.Schema:
				if e == nil {
					break
				}
				elementType = &valueType{Type: e.Type}
			default:
				log.Errorf(":ohno: Elem is something unexpected: %#v", e)
			}
		}

		// the default type for a map value is a string, so if
		// it hasn't been explicitly set, set it here.
		if s.Type == schema.TypeMap && elementType == nil {
			elementType = &valueType{Type: schema.TypeString}
		}

		def.Type = &valueType{
			Type:        s.Type,
			ElementType: elementType,
		}

		// also schema.Elem + others. see:
		// https://github.com/terraform-providers/terraform-provider-aws/blob/master/aws/data_source_aws_lb.go#L57

		msg.Fields = append(msg.Fields, def)
	}

	return msg, nil
}

type messageTemplateContext struct {
	Name     string
	Indent   string
	Messages []string
	Fields   []string
}

func (msg *messageDef) protoFields(parent *context) ([]string, error) {
	ctx := &context{parent, msg}

	fields := make([]string, 0, len(msg.Fields))

	for _, f := range msg.Fields {
		strField, err := f.protoField(ctx)
		if err != nil {
			log.Errorf("%s.%s: %s", msg.Name, f.Name, err)
			continue
		}
		fields = append(fields, strField)
	}

	return fields, nil
}

var messageTemplate = template.Must(template.New("message").Parse(`{{.Indent}}message {{.Name}} {{"{"}}{{range .Messages}}
{{.}}
{{end}}{{range .Fields}}
{{$.Indent}}  {{.}}{{end}}
{{.Indent}}}`))

func (msg *messageDef) string(indent string) (string, error) {
	var buf bytes.Buffer

	var submsgs []string
	for _, submsg := range msg.Messages {
		msgStr, err := submsg.string(indent + "  ")
		if err != nil {
			errors.Wrapf(err, "submsg.string")
		}
		submsgs = append(submsgs, msgStr)
	}

	strFields, err := msg.protoFields(nil)
	if err != nil {
		errors.Wrapf(err, "msg.protoFields")
	}
	tmplCtx := messageTemplateContext{
		Name:     msg.Name,
		Indent:   indent,
		Messages: submsgs,
		Fields:   strFields,
	}
	err = messageTemplate.Execute(&buf, &tmplCtx)
	if err != nil {
		errors.Wrapf(err, "messageTemplate.Execute")
	}

	return buf.String(), nil
}

func (msg *messageDef) String() string {
	str, err := msg.string("")
	if err != nil {
		panic(fmt.Sprintf("msg[%s].String(): %s", msg.Name, err))
	}

	return str
}

const (
	usage = `Usage: %s [OPTION...] [--] DIR
Emit protobufs for Terraform resources to a directory.
`
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
	}
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	dirPath := flag.Arg(0)

	err := os.MkdirAll(dirPath, 0775)
	if err != nil {
		log.Fatalf("MkdirAll(%q) failed: %s", dirPath, err)
	}

	tfProvider := aws.Provider()
	provider := tfProvider.(*schema.Provider)

	for name, resource := range provider.ResourcesMap {
		if !strings.HasPrefix(name, "aws_lb") {
			continue
		}

		if resource.Schema == nil {
			log.Warnf("resource %s missing schema", name)
			continue
		}

		name = strings.TrimPrefix(name, "aws_")

		msg, err := newMessage(name, resource)
		if err != nil {
			log.Errorf("newMessage(%s): %s", name, err)
		}

		path := fmt.Sprintf("%s/%s.proto", dirPath, name)

		contents := fmt.Sprintf(`syntax = "proto3";`+"\npackage terraform.aws;\nimport \"validate/validate.proto\";\n\n%s\n", msg.String())

		if err = ioutil.WriteFile(path, []byte(contents), 0644); err != nil {
			log.Fatalf("WriteFile(%q): %s", path, err)
		}
	}
}
