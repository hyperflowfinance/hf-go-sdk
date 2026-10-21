package codegen

import (
	"fmt"
	"strconv"
	"strings"
)

func objectStructName(name string, schema *Schema) string {
	if schema.SubscriptionType != nil && name == schema.SubscriptionType.Name {
		return goName(name) + "Event"
	}
	return goName(name)
}

func renderObject(directory string, typeDef Type, schema *Schema, types map[string]Type) error {
	name := objectStructName(typeDef.Name, schema)
	if len(typeDef.Fields) <= 55 {
		return renderInlineObject(directory, name, typeDef, schema, types)
	}
	groups := (len(typeDef.Fields) + 49) / 50
	var root strings.Builder
	root.WriteString(sourceHeader())
	root.WriteString("type " + name + " struct {\n")
	for group := 0; group < groups; group++ {
		root.WriteString("\t" + name + "Part" + strconv.Itoa(group) + "\n")
	}
	root.WriteString("}\n")
	if err := writeSource(directory, "zz_generated_object_"+fileToken(typeDef.Name)+".go", root.String()); err != nil {
		return err
	}
	for group := 0; group < groups; group++ {
		start := group * 50
		end := min(start+50, len(typeDef.Fields))
		if err := renderObjectGroup(directory, name, group, typeDef, typeDef.Fields[start:end], schema, types); err != nil {
			return err
		}
	}
	return nil
}

func renderInlineObject(directory, name string, typeDef Type, schema *Schema, types map[string]Type) error {
	lines, usesJSON, err := renderObjectFields(typeDef.Name, typeDef.Fields, schema, types)
	if err != nil {
		return err
	}
	var source strings.Builder
	source.WriteString(sourceHeader())
	if usesJSON {
		source.WriteString("import \"encoding/json\"\n\n")
	}
	source.WriteString("type " + name + " struct {\n" + lines + "}\n")
	return writeSource(directory, "zz_generated_object_"+fileToken(typeDef.Name)+".go", source.String())
}

func renderObjectGroup(directory, name string, group int, typeDef Type, fields []Field, schema *Schema, types map[string]Type) error {
	lines, usesJSON, err := renderObjectFields(typeDef.Name, fields, schema, types)
	if err != nil {
		return err
	}
	var source strings.Builder
	source.WriteString(sourceHeader())
	if usesJSON {
		source.WriteString("import \"encoding/json\"\n\n")
	}
	source.WriteString("type " + name + "Part" + strconv.Itoa(group) + " struct {\n" + lines + "}\n")
	file := fmt.Sprintf("zz_generated_object_%s_%d.go", fileToken(typeDef.Name), group)
	return writeSource(directory, file, source.String())
}
