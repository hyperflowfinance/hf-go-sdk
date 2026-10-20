package codegen

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
)

func writeSource(directory, name, source string) error {
	formatted, err := format.Source([]byte(source))
	if err != nil {
		return fmt.Errorf("format generated %s: %w", name, err)
	}
	if lines := strings.Count(string(formatted), "\n"); lines >= 100 {
		return fmt.Errorf("generated file %s has %d lines; limit is 99", name, lines)
	}
	if err := os.WriteFile(filepath.Join(directory, name), formatted, 0o644); err != nil {
		return fmt.Errorf("write generated %s: %w", name, err)
	}
	return nil
}

func sourceHeader() string {
	return "package hyperflow\n\n"
}
