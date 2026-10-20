package hyperflow

import (
	"fmt"
	"strings"
)

func renderFieldNode(field *selectionNode, counter *int, defs *[]string, vars map[string]any) (string, error) {
	args := make([]string, 0, len(field.args))
	for _, arg := range field.args {
		name := fmt.Sprintf("v%d", *counter)
		*counter++
		value := arg.value
		if arg.optional {
			optional, ok := value.(optionalValue)
			if !ok {
				return "", fmt.Errorf("optional argument %q has invalid value type", arg.name)
			}
			var set bool
			value, set = optional.graphOptional()
			if !set {
				continue
			}
		}
		*defs = append(*defs, "$"+name+": "+arg.gqlType)
		vars[name] = value
		args = append(args, arg.name+": $"+name)
	}
	text := field.name
	if len(args) != 0 {
		text += "(" + strings.Join(args, ", ") + ")"
	}
	if len(field.children) != 0 {
		body, err := renderFields(field.children, counter, defs, vars)
		if err != nil {
			return "", err
		}
		text += " { " + body + " }"
	} else if field.requiresSelection {
		return "", fmt.Errorf("field %q requires a selection set", field.name)
	}
	return text, nil
}
