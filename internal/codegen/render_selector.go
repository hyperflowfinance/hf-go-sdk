package codegen

import (
	"fmt"
	"strings"
)

func renderSelector(directory string, typeDef Type, types map[string]Type) error {
	name := goName(typeDef.Name) + "Fields"
	typeFile := "zz_generated_selector_" + fileToken(typeDef.Name) + ".go"
	if err := writeSource(directory, typeFile, sourceHeader()+"type "+name+" struct{ node *selectionNode }\n"); err != nil {
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
	for start, part := 0, 0; start < len(methods); start, part = start+12, part+1 {
		end := min(start+12, len(methods))
		var source strings.Builder
		source.WriteString(sourceHeader())
		for _, method := range methods[start:end] {
			source.WriteString(method + "\n")
		}
		file := fmt.Sprintf("zz_generated_selector_%s_%d.go", fileToken(typeDef.Name), part)
		if err := writeSource(directory, file, source.String()); err != nil {
			return err
		}
	}
	return nil
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
