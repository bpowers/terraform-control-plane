package main

import (
	"bytes"
	"fmt"
	"sort"
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
	protoType(ctx *context) (string, error)
}

type valueType struct {
	Type        schema.ValueType
	ElementType fieldType
}

func (ty *valueType) protoType(ctx *context) (string, error) {
	var err error
	ptype := ""

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
		log.Infof("TODO: add gogoprotobuf unique constraint for Set")
		fallthrough
	case schema.TypeList:
		ptype = "repeated "
		// TODO fix this after fixing resources below
		if ty.ElementType != nil {
			eType, err := ty.ElementType.protoType(ctx)
			if err != nil {
				return "", errors.Wrapf(err, "%#v.protoType", ty.ElementType)
			}
			ptype += eType
		} else {
			log.Warnf("nil repeated type for %#v", ty)
		}
	case schema.TypeMap:
		var valueType string
		if ty.ElementType != nil {
			valueType, err = ty.ElementType.protoType(ctx)
			if err != nil {
				return "", errors.Wrapf(err, "%#v.protoType", ty.ElementType)
			}
		} else {
			log.Warnf("nil map type for %#v", ty)
		}
		ptype = fmt.Sprintf("map<string, %s>", valueType)
	default:
		return "", errors.Errorf(":ohno: unknown Schema.Type %#v", ty.Type)
	}

	return ptype, nil
}

type referenceType struct {
	Name string
}

func (ty *referenceType) protoType(ctx *context) (string, error) {
	log.Warnf("TODO: reference type %s", ty.Name)
	return "/shrug", nil
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

	ty, err := f.Type.protoType(ctx)
	if err != nil {
		return "", err
	}
	name := strcase.ToSnake(f.Name)
	return fmt.Sprintf("%s %s = %d;", ty, name, f.Number), nil
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
			Number:   uint(i),
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
				log.Infof("TODO: field %s elem resource: %v", name, e)
			case *schema.Schema:
				if e == nil {
					break
				}
				elementType = &valueType{Type: e.Type}
			}
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

var messageTemplate = template.Must(template.New("message").Parse(`{{.Indent}}message {{.Name}} {
{{range .Messages}}
{{$.Indent}}  {{.}}{{end}}
{{range .Fields}}
{{$.Indent}}  {{.}}{{end}}
{{.Indent}}}`))

func (msg *messageDef) string(indent string) (string, error) {
	var buf bytes.Buffer

	strFields, err := msg.protoFields(nil)
	if err != nil {
		errors.Wrapf(err, "msg.protoFields")
	}
	tmplCtx := messageTemplateContext{
		Name:   msg.Name,
		Indent: indent,
		Fields: strFields,
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

func main() {
	tfProvider := aws.Provider()
	provider := tfProvider.(*schema.Provider)

	for name, resource := range provider.ResourcesMap {
		if name != "aws_lb" {
			continue
		}

		if resource.Schema == nil {
			log.Warnf("resource %s missing schema", name)
			continue
		}

		msg, err := newMessage(name, resource)
		if err != nil {
			log.Errorf("newMessage(%s): %s", name, err)
		}

		fmt.Printf("%s\n", msg)
	}
}
