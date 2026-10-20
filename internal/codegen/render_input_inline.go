package codegen

import (
	"fmt"
	"strconv"
	"strings"
)

func renderInlineInput(directory, name string, typeDef Type, types map[string]Type) error {
	var source strings.Builder
	source.WriteString(sourceHeader())
	if needsJSONInput(typeDef.InputFields, types) {
		source.WriteString("import \"encoding/json\"\n\n")
	}
	source.WriteString("type " + name + " struct {\n")
	for _, field := range typeDef.InputFields {
		fieldType, err := inputType(field.Type, types)
		if err != nil {
			return fmt.Errorf("input %s.%s: %w", typeDef.Name, field.Name, err)
		}
		source.WriteString("\t" + goName(field.Name) + " " + fieldType + " `json:" + strconv.Quote(field.Name) + "`\n")
	}
	source.WriteString("}\n\nfunc (v " + name + ") MarshalJSON() ([]byte, error) {\n")
	source.WriteString("\treturn marshalInput(v)\n}\n")
	return writeSource(directory, "zz_generated_input_"+fileToken(typeDef.Name)+".go", source.String())
}

func needsJSONInput(fields []Input, types map[string]Type) bool {
	for _, field := range fields {
		fieldType, err := inputType(field.Type, types)
		if err == nil && strings.Contains(fieldType, "json.RawMessage") {
			return true
		}
	}
	return false
}
