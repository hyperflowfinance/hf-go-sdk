package codegen

import (
	"fmt"
	"strconv"
	"strings"
)

func renderField(receiver, method string, field Field, types map[string]Type) (string, error) {
	params, args, err := renderArguments(field.Args, types)
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
	argsText := renderArgumentList(args)
	var source strings.Builder
	if composite {
		childType := goName(target) + "Fields"
		params = append(params, "selectFields func(*"+childType+")")
		source.WriteString("func (s *" + receiver + ") " + method + "(" + strings.Join(params, ", ") + ") *" + receiver + " {\n")
		source.WriteString("\tchild := s.node.object(" + strconv.Quote(field.Name) + argsText + ")\n")
		source.WriteString("\tif selectFields != nil { selectFields(&" + childType + "{node: child}) }\n")
	} else {
		source.WriteString("func (s *" + receiver + ") " + method + "(" + strings.Join(params, ", ") + ") *" + receiver + " {\n")
		source.WriteString("\ts.node.field(" + strconv.Quote(field.Name) + argsText + ")\n")
	}
	source.WriteString("\treturn s\n}\n")
	return source.String(), nil
}

func renderArguments(inputs []Input, types map[string]Type) ([]string, []string, error) {
	params := make([]string, 0, len(inputs))
	args := make([]string, 0, len(inputs))
	used := map[string]bool{}
	for _, input := range inputs {
		name := "arg" + goName(input.Name)
		if used[name] {
			return nil, nil, fmt.Errorf("arguments collide as Go name %q", name)
		}
		used[name] = true
		typeName, err := inputType(input.Type, types)
		if err != nil {
			return nil, nil, err
		}
		gqlName, err := gqlType(input.Type)
		if err != nil {
			return nil, nil, err
		}
		optional := input.Type.Kind != "NON_NULL"
		params = append(params, name+" "+typeName)
		args = append(args, "argument{name: "+strconv.Quote(input.Name)+", gqlType: "+strconv.Quote(gqlName)+
			", value: "+name+", optional: "+strconv.FormatBool(optional)+"}")
	}
	return params, args, nil
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
