package codegen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRenderSelectorSplitsByFormattedLineCount(t *testing.T) {
	query := Type{Kind: "OBJECT", Name: "Query"}
	for i := 0; i < 20; i++ {
		query.Fields = append(query.Fields, Field{
			Name: fmt.Sprintf("item%d", i),
			Type: TypeRef{Kind: "OBJECT", Name: "Item"},
			Args: []Input{
				{Name: "limit", Type: TypeRef{Kind: "SCALAR", Name: "Int"}},
				{Name: "cursor", Type: TypeRef{Kind: "SCALAR", Name: "JSON"}},
			},
		})
	}
	types := map[string]Type{
		"Item": {Kind: "OBJECT", Name: "Item"},
		"JSON": {Kind: "SCALAR", Name: "JSON"},
	}
	output := t.TempDir()
	if err := renderSelector(output, query, types); err != nil {
		t.Fatal(err)
	}
	files, err := filepath.Glob(filepath.Join(output, "zz_generated_selector_query_*.go"))
	if err != nil {
		t.Fatal(err)
	}
	if len(files) < 3 {
		t.Fatalf("generated %d files, want at least 3", len(files))
	}
	var selectors strings.Builder
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		source := string(data)
		if lines := strings.Count(source, "\n"); lines >= 100 {
			t.Fatalf("%s has %d lines", file, lines)
		}
		if strings.Contains(source, "json.RawMessage") && !strings.Contains(source, `import "encoding/json"`) {
			t.Fatalf("%s is missing its JSON import", file)
		}
		selectors.WriteString(source)
	}
	source := selectors.String()
	if count := strings.Count(source, "func (s *QueryFields) Typename("); count != 1 {
		t.Errorf("Typename generated %d times, want 1", count)
	}
	for _, field := range query.Fields {
		method := goName(field.Name)
		for _, declaration := range []string{
			"func (s *QueryFields) " + method + "(",
			"type Query" + method + "Args struct",
		} {
			if count := strings.Count(source, declaration); count != 1 {
				t.Errorf("%q generated %d times, want 1", declaration, count)
			}
		}
	}
}

func TestRenderSelectorRejectsOversizedMethod(t *testing.T) {
	field := Field{Name: "item", Type: TypeRef{Kind: "OBJECT", Name: "Item"}}
	for i := 0; i < 100; i++ {
		field.Args = append(field.Args, Input{
			Name: fmt.Sprintf("arg%d", i),
			Type: TypeRef{Kind: "SCALAR", Name: "Int"},
		})
	}
	err := renderSelector(t.TempDir(), Type{
		Kind: "OBJECT", Name: "Query", Fields: []Field{field},
	}, map[string]Type{"Item": {Kind: "OBJECT", Name: "Item"}})
	if err == nil || !strings.Contains(err.Error(), "limit is 99") {
		t.Fatalf("got %v, want explicit line-limit error", err)
	}
}
