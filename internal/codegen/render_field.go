package codegen

import (
	"fmt"
	"strconv"
	"strings"
)

func renderField(ownerType, receiver, method string, field Field, types map[string]Type) (string, error) {
	argsType := ownerType + method + "Args"
	decl, params, args, hasOpt, err := renderArguments(argsType, field.Args, types)
	if err != nil {
		return "", err
	}
	target, err := namedRef(field.Type)
	if err != nil {
		return "", err
	}
	typeDef := types[target]
	if typeDef.Name == "" {
		return "", fmt.Errorf("return type %q is missing from introspection", target)
	}
	composite := typeDef.Kind == "OBJECT" || typeDef.Kind == "INTERFACE" || typeDef.Kind == "UNION"
	if !composite && typeDef.Kind != "SCALAR" && typeDef.Kind != "ENUM" {
		return "", fmt.Errorf("return type %q has unsupported kind %q", target, typeDef.Kind)
	}
	if composite {
		params = append(params, "selectFields func(*"+goName(target)+"Fields)")
	}
	if hasOpt {
		params = append(params, "opts ..."+argsType)
	}
	var source strings.Builder
	source.WriteString(decl)
	source.WriteString("func (s *" + receiver + ") " + method + "(" + strings.Join(params, ", ") + ") *" + receiver + " {\n")
	if hasOpt {
		source.WriteString("\tvar opt " + argsType + "\n\tif len(opts) > 0 { opt = opts[0] }\n")
	}
	argsText := renderArgumentList(args)
	if composite {
		childType := goName(target) + "Fields"
		source.WriteString("\tchild := s.node.object(" + strconv.Quote(field.Name) + argsText + ")\n")
		source.WriteString("\tif selectFields != nil { selectFields(&" + childType + "{node: child}) }\n")
	} else {
		source.WriteString("\ts.node.field(" + strconv.Quote(field.Name) + argsText + ")\n")
	}
	source.WriteString("\treturn s\n}\n")
	return source.String(), nil
}

func renderArguments(argsType string, inputs []Input, types map[string]Type) (string, []string, []string, bool, error) {
	var params, args, optFields []string
	used := map[string]bool{}
	for _, input := range inputs {
		fieldName := goName(input.Name)
		if used[fieldName] {
			return "", nil, nil, false, fmt.Errorf("arguments collide as Go name %q", fieldName)
		}
		used[fieldName] = true
		typeName, err := inputType(input.Type, types)
		if err != nil {
			return "", nil, nil, false, err
		}
		gqlName, err := gqlType(input.Type)
		if err != nil {
			return "", nil, nil, false, err
		}
		optional := input.Type.Kind != "NON_NULL"
		valRef := "arg" + fieldName
		if optional {
			optFields = append(optFields, "\t"+fieldName+" "+typeName)
			valRef = "opt." + fieldName
		} else {
			params = append(params, valRef+" "+typeName)
		}
		args = append(args, "argument{name: "+strconv.Quote(input.Name)+", gqlType: "+strconv.Quote(gqlName)+
			", value: "+valRef+", optional: "+strconv.FormatBool(optional)+"}")
	}
	var decl string
	if len(optFields) > 0 {
		decl = "type " + argsType + " struct {\n" + strings.Join(optFields, "\n") + "\n}\n\n"
	}
	return decl, params, args, len(optFields) > 0, nil
}

func renderArgumentList(args []string) string {
	if len(args) == 0 {
		return ""
	}
	return ", " + strings.Join(args, ", ")
}

func scalarMethod(receiver, field, method string) string {
	return "func (s *" + receiver + ") " + method + "() *" + receiver + " {\n" +
		"\ts.node.field(" + strconv.Quote(field) + ")\n\treturn s\n}\n"
}
