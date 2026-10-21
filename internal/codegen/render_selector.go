package codegen

import (
	"fmt"
	"strings"
)

func renderSelector(directory string, typeDef Type, types map[string]Type) error {
	name := goName(typeDef.Name) + "Fields"
	typeFile := "zz_generated_selector_" + fileToken(typeDef.Name) + ".go"
	typeSource := sourceHeader() + "// " + name + " builds a selection set for GraphQL " + typeDef.Name + ".\n" +
		"type " + name + " struct{ node *selectionNode }\n"
	if err := writeSource(directory, typeFile, typeSource); err != nil {
		return err
	}
	methods := []string{scalarMethod(name, "__typename", "Typename")}
	methodNames := []string{"Typename"}
	for _, field := range typeDef.Fields {
		methodName := goName(field.Name)
		methodNames = append(methodNames, methodName)
		method, err := renderField(name, methodName, field, types)
		if err != nil {
			return fmt.Errorf("selector %s.%s: %w", typeDef.Name, field.Name, err)
		}
		methods = append(methods, method)
	}
	for _, possible := range typeDef.PossibleTypes {
		methodName := "On" + goName(possible.Name)
		methodNames = append(methodNames, methodName)
		method, err := renderFragment(name, methodName, possible.Name, types)
		if err != nil {
			return fmt.Errorf("selector %s fragment %s: %w", typeDef.Name, possible.Name, err)
		}
		methods = append(methods, method)
	}
	if err := uniqueGoNames(typeDef.Name, methodNames); err != nil {
		return err
	}
	var source strings.Builder
	source.WriteString(sourceHeader())
	part, lines := 0, 2
	for _, method := range methods {
		methodLines := strings.Count(method, "\n") + 1
		if lines+methodLines >= 75 {
			if err := writeSelectorPart(directory, typeDef.Name, part, source.String()); err != nil {
				return err
			}
			part++
			source.Reset()
			source.WriteString(sourceHeader())
			lines = 2
		}
		source.WriteString(method + "\n")
		lines += methodLines
	}
	return writeSelectorPart(directory, typeDef.Name, part, source.String())
}

func writeSelectorPart(directory, typeName string, part int, source string) error {
	file := fmt.Sprintf("zz_generated_selector_%s_%d.go", fileToken(typeName), part)
	return writeSource(directory, file, source)
}

func uniqueGoNames(typeName string, names []string) error {
	seen := map[string]bool{}
	for _, name := range names {
		if seen[name] {
			return fmt.Errorf("type %s has fields that collide as Go method %q", typeName, name)
		}
		seen[name] = true
	}
	return nil
}
