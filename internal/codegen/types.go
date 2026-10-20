package codegen

import "fmt"

func gqlType(ref TypeRef) (string, error) {
	switch ref.Kind {
	case "NON_NULL":
		inner, err := gqlTypeRef(ref.OfType)
		return inner + "!", err
	case "LIST":
		inner, err := gqlTypeRef(ref.OfType)
		return "[" + inner + "]", err
	case "SCALAR", "ENUM", "INPUT_OBJECT", "OBJECT", "INTERFACE", "UNION":
		if ref.Name == "" {
			return "", fmt.Errorf("named GraphQL type has no name")
		}
		return ref.Name, nil
	default:
		return "", fmt.Errorf("unsupported GraphQL type kind %q", ref.Kind)
	}
}

func gqlTypeRef(ref *TypeRef) (string, error) {
	if ref == nil {
		return "", fmt.Errorf("GraphQL type reference is incomplete")
	}
	return gqlType(*ref)
}

func namedRef(ref TypeRef) (string, error) {
	for ref.Kind == "NON_NULL" || ref.Kind == "LIST" {
		if ref.OfType == nil {
			return "", fmt.Errorf("GraphQL type reference is incomplete")
		}
		ref = *ref.OfType
	}
	if ref.Name == "" {
		return "", fmt.Errorf("GraphQL named type is missing")
	}
	return ref.Name, nil
}
