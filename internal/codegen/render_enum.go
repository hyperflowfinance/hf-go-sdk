package codegen

import (
	"fmt"
	"strconv"
	"strings"
)

func renderEnum(directory string, typeDef Type) error {
	name := goName(typeDef.Name)
	file := "zz_generated_enum_" + fileToken(typeDef.Name) + "_type.go"
	if err := writeSource(directory, file, sourceHeader()+"type "+name+" string\n"); err != nil {
		return err
	}
	values := make([]string, 0, len(typeDef.EnumValues))
	seen := map[string]string{}
	for _, value := range typeDef.EnumValues {
		constant := name + goName(value.Name)
		if previous, ok := seen[constant]; ok {
			return fmt.Errorf("enum %s values %q and %q map to %s", name, previous, value.Name, constant)
		}
		seen[constant] = value.Name
		values = append(values, constant+" "+name+" = "+strconv.Quote(value.Name))
	}
	for start, part := 0, 0; start < len(values); start, part = start+50, part+1 {
		end := min(start+50, len(values))
		var source strings.Builder
		source.WriteString(sourceHeader())
		source.WriteString("const (\n")
		for _, value := range values[start:end] {
			source.WriteString(value + "\n")
		}
		source.WriteString(")\n")
		if err := writeSource(directory, fmt.Sprintf("zz_generated_enum_%s_%d.go", fileToken(typeDef.Name), part), source.String()); err != nil {
			return err
		}
	}
	return nil
}

func fileToken(name string) string {
	return strings.ToLower(goName(name))
}
