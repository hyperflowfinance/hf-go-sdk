package codegen

import (
	"fmt"
	"strconv"
	"strings"
)

func renderInput(directory string, typeDef Type, types map[string]Type) error {
	name := inputName(typeDef.Name)
	fields := make([]string, 0, len(typeDef.InputFields))
	for _, field := range typeDef.InputFields {
		fields = append(fields, field.Name)
	}
	if err := uniqueNamesFrom(fields); err != nil {
		return fmt.Errorf("input %s: %w", typeDef.Name, err)
	}
	if len(typeDef.InputFields) <= 55 {
		return renderInlineInput(directory, name, typeDef, types)
	}
	groups := (len(typeDef.InputFields) + 49) / 50
	var root strings.Builder
	root.WriteString(sourceHeader())
	root.WriteString("// " + name + " contains values for GraphQL " + typeDef.Name + ".\n")
	root.WriteString("type " + name + " struct {\n")
	for group := 0; group < groups; group++ {
		root.WriteString(name + "Fields" + strconv.Itoa(group) + "\n")
	}
	root.WriteString("}\n\nfunc (v " + name + ") MarshalJSON() ([]byte, error) {\n")
	root.WriteString("\treturn marshalInput(v)\n}\n")
	if err := writeSource(directory, "zz_generated_input_"+fileToken(typeDef.Name)+".go", root.String()); err != nil {
		return err
	}
	for group := 0; group < groups; group++ {
		start := group * 50
		end := min(start+50, len(typeDef.InputFields))
		var source strings.Builder
		source.WriteString(sourceHeader())
		source.WriteString("// " + name + "Fields" + strconv.Itoa(group) + " contains part of " + name + ".\n")
		for _, field := range typeDef.InputFields[start:end] {
			fieldType, err := inputType(field.Type, types)
			if err != nil {
				return fmt.Errorf("input %s.%s: %w", typeDef.Name, field.Name, err)
			}
			if strings.Contains(fieldType, "json.RawMessage") {
				source.WriteString("import \"encoding/json\"\n\n")
				break
			}
		}
		source.WriteString("type " + name + "Fields" + strconv.Itoa(group) + " struct {\n")
		for _, field := range typeDef.InputFields[start:end] {
			fieldType, err := inputType(field.Type, types)
			if err != nil {
				return fmt.Errorf("input %s.%s: %w", typeDef.Name, field.Name, err)
			}
			source.WriteString("\t" + goName(field.Name) + " " + fieldType + " `json:" + strconv.Quote(field.Name) + "`\n")
		}
		source.WriteString("}\n")
		if err := writeSource(directory, fmt.Sprintf("zz_generated_input_%s_%d.go", fileToken(typeDef.Name), group), source.String()); err != nil {
			return err
		}
	}
	return nil
}

func uniqueNamesFrom(names []string) error {
	seen := map[string]string{}
	for _, name := range names {
		goName := goName(name)
		if previous, ok := seen[goName]; ok {
			return fmt.Errorf("fields %q and %q map to the same Go name %q", previous, name, goName)
		}
		seen[goName] = name
	}
	return nil
}
