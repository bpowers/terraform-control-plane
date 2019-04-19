package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

func main() {
	tfProvider := aws.Provider()
	provider := tfProvider.(*schema.Provider)

	for name, resource := range provider.ResourcesMap {
		if name == "aws_lb" {
			fmt.Printf("%s: %v\n", name, resource.SchemaVersion)
			for field, schema := range resource.Schema {
				fmt.Printf("\t%s: %v (%v)\n", field, schema.Type, schema.Optional)
				// also schema.Elem + others. see:
				// https://github.com/terraform-providers/terraform-provider-aws/blob/master/aws/data_source_aws_lb.go#L57
			}
		}
	}
}
