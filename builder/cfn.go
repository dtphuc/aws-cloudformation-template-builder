package builder

type CfnBuilder struct {
	Builder
}

func NewCfnBuilder(includeOptional, buildIamPolicies bool) CfnBuilder {
	var b CfnBuilder
	b.Spec = CfnSpec
	b.IncludeOptionalProperties = includeOptional
	b.BuildIamPolicies = buildIamPolicies

	return b
}

// Template produces a CloudFormation template for the
// resources in the config map
func (b CfnBuilder) Template(config map[string]string) (map[string]interface{}, map[interface{}]interface{}) {
	// Generate resources
	resources := make(map[string]interface{})
	comments := make(map[string]interface{})
	for name, resourceType := range config {
		resources[name], comments[name] = b.newResource(resourceType)
	}

	// Build the template
	return map[string]interface{}{
			"AWSTemplateFormatVersion": "2010-09-09",
			"Description":              "Template generated by cfn-skeleton",
			"Resources":                resources,
			// TODO: "Outputs": outputs,
		}, map[interface{}]interface{}{
			"Resources": comments,
		}
}
