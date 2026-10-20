package codegen

import (
	"fmt"
	"strings"
)

func validateNames(schema *Schema, types map[string]Type) error {
	generated := map[string]string{
		"Client": "runtime", "Config": "runtime", "Document": "runtime",
		"GraphQLError": "runtime", "GraphQLErrors": "runtime", "HTTPError": "runtime",
		"Optional": "runtime", "Some": "runtime", "Connect": "runtime",
		"NewQuery": "query root", "NewMutation": "mutation root", "NewSubscription": "subscription root",
	}
	for _, typeDef := range schema.Types {
		if strings.HasPrefix(typeDef.Name, "__") {
			continue
		}
		name := generatedTypeName(typeDef)
		if name == "" {
			continue
		}
		if previous, exists := generated[name]; exists {
			return fmt.Errorf("GraphQL type %q collides with generated %s name %q", typeDef.Name, previous, name)
		}
		generated[name] = typeDef.Name
	}
	for _, typeDef := range schema.Types {
		if typeDef.Kind != "ENUM" || strings.HasPrefix(typeDef.Name, "__") {
			continue
		}
		for _, value := range typeDef.EnumValues {
			name := goName(typeDef.Name) + goName(value.Name)
			if previous, exists := generated[name]; exists {
				return fmt.Errorf("enum value %q.%s collides with generated %s name %q", typeDef.Name, value.Name, previous, name)
			}
			generated[name] = "enum value " + typeDef.Name + "." + value.Name
		}
	}
	for _, root := range []*RootType{schema.QueryType, schema.MutationType, schema.SubscriptionType} {
		if root != nil {
			if _, ok := types[root.Name]; !ok {
				return fmt.Errorf("root type %q is absent from introspection types", root.Name)
			}
		}
	}
	return nil
}

func generatedTypeName(typeDef Type) string {
	switch typeDef.Kind {
	case "ENUM":
		return goName(typeDef.Name)
	case "INPUT_OBJECT":
		return inputName(typeDef.Name)
	case "OBJECT", "INTERFACE", "UNION":
		return goName(typeDef.Name) + "Fields"
	default:
		return ""
	}
}
