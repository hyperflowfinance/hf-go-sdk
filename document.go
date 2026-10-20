package hyperflow

import (
	"fmt"
	"strings"
)

func (d Document) render() (string, map[string]any, error) {
	if d.root == nil || len(d.root.children) == 0 {
		return "", nil, fmt.Errorf("GraphQL operation has no selected fields")
	}
	variables := make(map[string]any)
	definitions := []string{}
	counter := 0
	body, err := renderFields(d.root.children, &counter, &definitions, variables)
	if err != nil {
		return "", nil, err
	}
	header := d.operation
	if len(definitions) != 0 {
		header += "(" + strings.Join(definitions, ", ") + ")"
	}
	return header + " { " + body + " }", variables, nil
}

func renderFields(fields []*selectionNode, counter *int, defs *[]string, vars map[string]any) (string, error) {
	out := make([]string, 0, len(fields))
	for _, field := range fields {
		if field.fragment != "" {
			if len(field.children) == 0 {
				return "", fmt.Errorf("inline fragment on %q requires a selection set", field.fragment)
			}
			body, err := renderFields(field.children, counter, defs, vars)
			if err != nil {
				return "", err
			}
			out = append(out, "... on "+field.fragment+" { "+body+" }")
			continue
		}
		text, err := renderFieldNode(field, counter, defs, vars)
		if err != nil {
			return "", err
		}
		out = append(out, text)
	}
	return strings.Join(out, " "), nil
}
