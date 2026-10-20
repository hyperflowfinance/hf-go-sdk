package codegen

import "fmt"

func inputType(ref TypeRef, types map[string]Type) (string, error) {
	if ref.Kind != "NON_NULL" {
		base, err := nonNullInputType(ref, types)
		return "Optional[" + base + "]", err
	}
	if ref.OfType == nil {
		return "", fmt.Errorf("non-null GraphQL type reference is incomplete")
	}
	return nonNullInputType(*ref.OfType, types)
}

func nonNullInputType(ref TypeRef, types map[string]Type) (string, error) {
	if ref.Kind == "LIST" {
		if ref.OfType == nil {
			return "", fmt.Errorf("list item type is missing")
		}
		item, err := listItemType(*ref.OfType, types)
		return "[]" + item, err
	}
	return namedInputType(ref.Name, types)
}

func listItemType(ref TypeRef, types map[string]Type) (string, error) {
	if ref.Kind == "NON_NULL" {
		if ref.OfType == nil {
			return "", fmt.Errorf("non-null list item type is incomplete")
		}
		return nonNullInputType(*ref.OfType, types)
	}
	base, err := nonNullInputType(ref, types)
	return "*" + base, err
}

func namedInputType(name string, types map[string]Type) (string, error) {
	switch name {
	case "String", "ID":
		return "string", nil
	case "Boolean":
		return "bool", nil
	case "Int":
		return "int32", nil
	case "Float":
		return "float64", nil
	}
	typeDef, ok := types[name]
	if !ok {
		return "", fmt.Errorf("GraphQL type %q is not in introspection schema", name)
	}
	switch typeDef.Kind {
	case "ENUM":
		return goName(name), nil
	case "INPUT_OBJECT":
		return inputName(name), nil
	case "SCALAR":
		return "json.RawMessage", nil
	default:
		return "", fmt.Errorf("GraphQL type %q cannot be used as input", name)
	}
}
