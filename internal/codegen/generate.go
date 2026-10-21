package codegen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Generate(schema *Schema, directory string) error {
	if schema == nil || schema.QueryType == nil {
		return fmt.Errorf("schema has no query root")
	}
	types := make(map[string]Type, len(schema.Types))
	for _, typeDef := range schema.Types {
		if typeDef.Name != "" {
			if _, exists := types[typeDef.Name]; exists {
				return fmt.Errorf("introspection schema repeats type %q", typeDef.Name)
			}
			types[typeDef.Name] = typeDef
		}
	}
	if err := validateNames(schema, types); err != nil {
		return err
	}
	if err := os.MkdirAll(directory, 0o755); err != nil {
		return fmt.Errorf("create generated-code directory: %w", err)
	}
	oldFiles, err := filepath.Glob(filepath.Join(directory, "zz_generated_*.go"))
	if err != nil {
		return fmt.Errorf("find stale generated files: %w", err)
	}
	for _, file := range oldFiles {
		if err := os.Remove(file); err != nil {
			return fmt.Errorf("remove stale generated file %s: %w", file, err)
		}
	}
	for _, typeDef := range schema.Types {
		if strings.HasPrefix(typeDef.Name, "__") {
			continue
		}
		switch typeDef.Kind {
		case "ENUM":
			if err := renderEnum(directory, typeDef); err != nil {
				return err
			}
		case "INPUT_OBJECT":
			if err := renderInput(directory, typeDef, types); err != nil {
				return err
			}
		case "OBJECT":
			if err := renderObject(directory, typeDef, schema, types); err != nil {
				return err
			}
			if err := renderSelector(directory, typeDef, types); err != nil {
				return err
			}
		case "INTERFACE", "UNION":
			if err := renderSelector(directory, typeDef, types); err != nil {
				return err
			}
		}
	}
	return renderOperations(directory, schema, types)
}
