package main

import (
	"bytes"
	"fmt"
	"sort"
	"text/template"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/iancoleman/strcase"
	log "github.com/sirupsen/logrus"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

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
	Name        string
	Number      uint
	Type        schema.ValueType
	Optional    bool
	ElementType *schema.ValueType
	MinItems    *uint
	MaxItems    *uint
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
			Type:     s.Type,
			Optional: s.Optional,
		}

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
				ty := e.Type
				def.ElementType = &ty
			}
		}
		// also schema.Elem + others. see:
		// https://github.com/terraform-providers/terraform-provider-aws/blob/master/aws/data_source_aws_lb.go#L57

		msg.Fields = append(msg.Fields, def)
	}

	return msg, nil
}

var messageTemplate = template.Must(template.New("message").Parse(`message {{.Name}} {
  
}`))

func (msg *messageDef) String() string {
	var buf bytes.Buffer
	err := messageTemplate.Execute(&buf, msg)
	if err != nil {
		panic("template execution failed!")
	}

	return buf.String()
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

		fmt.Printf("MSG: %s\n", msg)
	}
}
