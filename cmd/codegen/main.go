package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/hyperflowfinance/hf-go-sdk/internal/codegen"
)

func main() {
	schemaFile := flag.String("schema-file", "", "read an introspection response from a file")
	outputDir := flag.String("out", ".", "directory for generated Go files")
	flag.Parse()
	if err := run(*schemaFile, *outputDir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(schemaFile, outputDir string) error {
	var schema *codegen.Schema
	if schemaFile == "" {
		var err error
		schema, err = codegen.FetchSchema(context.Background())
		if err != nil {
			return err
		}
	} else {
		file, err := os.Open(schemaFile)
		if err != nil {
			return fmt.Errorf("open schema file: %w", err)
		}
		defer file.Close()
		schema, err = codegen.ParseSchema(file)
		if err != nil {
			return err
		}
	}
	return codegen.Generate(schema, outputDir)
}
