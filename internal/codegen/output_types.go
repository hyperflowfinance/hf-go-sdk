package codegen

import (
	"fmt"
	"strconv"
	"strings"
)

func renderObjectFields(typeName string, fields []Field, schema *Schema, types map[string]Type) (string, bool, error) {
	var out strings.Builder
	usesJSON := false
	for _, field := range fields {
		fieldType, err := outputType(field.Type, schema, types)
		if err != nil {
			return "", false, fmt.Errorf("object %s.%s: %w", typeName, field.Name, err)
		}
		if strings.Contains(fieldType, "json.RawMessage") {
			usesJSON = true
		}
		out.WriteString("\t" + goName(field.Name) + " " + fieldType + " `json:" + strconv.Quote(field.Name) + "`\n")
	}
	return out.String(), usesJSON, nil
}

func outputType(ref TypeRef, schema *Schema, types map[string]Type) (string, error) {
	if ref.Kind == "NON_NULL" {
		if ref.OfType == nil {
			return "", fmt.Errorf("non-null GraphQL type reference is incomplete")
		}
		return nonNullOutputType(*ref.OfType, schema, types)
	}
	base, err := nonNullOutputType(ref, schema, types)
	if err != nil || strings.HasPrefix(base, "[]") || base == "json.RawMessage" {
		return base, err
	}
	return "*" + base, nil
}

func nonNullOutputType(ref TypeRef, schema *Schema, types map[string]Type) (string, error) {
	if ref.Kind == "LIST" {
		if ref.OfType == nil {
			return "", fmt.Errorf("list item type is missing")
		}
		item, err := outputType(*ref.OfType, schema, types)
		return "[]" + item, err
	}
	switch ref.Name {
	case "String", "ID":
		return "string", nil
	case "Boolean":
		return "bool", nil
	case "Int":
		return "int32", nil
	case "Float":
		return "float64", nil
	}
	typeDef, ok := types[ref.Name]
	if !ok {
		return "", fmt.Errorf("GraphQL type %q is not in introspection schema", ref.Name)
	}
	switch typeDef.Kind {
	case "ENUM":
		return goName(ref.Name), nil
	case "OBJECT":
		return objectStructName(ref.Name, schema), nil
	case "SCALAR", "INTERFACE", "UNION":
		return "json.RawMessage", nil
	default:
		return "", fmt.Errorf("GraphQL type %q cannot be used as output", ref.Name)
	}
}
