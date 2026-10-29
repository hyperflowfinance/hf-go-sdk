package codegen

import (
	"fmt"
	"go/format"
	"strings"
)

func renderSelector(directory string, typeDef Type, types map[string]Type) error {
	typeName := goName(typeDef.Name)
	name := typeName + "Fields"
	typeFile := "zz_generated_selector_" + fileToken(typeDef.Name) + ".go"
	if err := writeSource(directory, typeFile, sourceHeader()+"type "+name+" struct{ node *selectionNode }\n"); err != nil {
		return err
	}
	methods := []string{scalarMethod(name, "__typename", "Typename")}
	methodNames := []string{"Typename"}
	for _, field := range typeDef.Fields {
		methodName := goName(field.Name)
		methodNames = append(methodNames, methodName)
		method, err := renderField(typeName, name, methodName, field, types)
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
	for start, part := 0, 0; start < len(methods); part++ {
		end := start
		var source string
		for end < len(methods) {
			body := strings.Join(methods[start:end+1], "\n") + "\n"
			candidate := sourceHeader()
			if strings.Contains(body, "json.RawMessage") {
				candidate += "import \"encoding/json\"\n\n"
			}
			formatted, err := format.Source([]byte(candidate + body))
			if err != nil {
				return fmt.Errorf("format selector %s: %w", typeDef.Name, err)
			}
			if strings.Count(string(formatted), "\n") >= 100 && end > start {
				break
			}
			source = string(formatted)
			end++
		}
		file := fmt.Sprintf("zz_generated_selector_%s_%d.go", fileToken(typeDef.Name), part)
		if err := writeSource(directory, file, source); err != nil {
			return err
		}
		start = end
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
