package codegen

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateFromIntrospectionFixture(t *testing.T) {
	file, err := os.Open("testdata/schema.json")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	schema, err := ParseSchema(file)
	if err != nil {
		t.Fatal(err)
	}
	output := t.TempDir()
	if err := Generate(schema, output); err != nil {
		t.Fatal(err)
	}
	checkGenerated(t, output, "zz_generated_selector_query_0.go", "Metadata")
	checkGenerated(t, output, "zz_generated_selector_metadata.go", "MetadataFields builds a selection set")
	checkGenerated(t, output, "zz_generated_selector_metadata_0.go", "// ChainID selects the GraphQL chainId field.")
	checkGenerated(t, output, "zz_generated_selector_searchresult_0.go", "OnThing", "OnUser")
	checkGenerated(t, output, "zz_generated_operations.go", "// NewQuery creates a query document", "NewMutation")
	for _, name := range generatedFiles(t, output) {
		data, err := os.ReadFile(filepath.Join(output, name))
		if err != nil {
			t.Fatal(err)
		}
		if lines := strings.Count(string(data), "\n"); lines >= 100 {
			t.Fatalf("%s has %d lines", name, lines)
		}
	}
}

func TestGenerateSubscriptionOperation(t *testing.T) {
	schema := &Schema{
		QueryType:        &RootType{Name: "Query"},
		SubscriptionType: &RootType{Name: "Subscription"},
		Types: []Type{
			{Kind: "OBJECT", Name: "Query"},
			{Kind: "OBJECT", Name: "Subscription"},
		},
	}
	output := t.TempDir()
	if err := Generate(schema, output); err != nil {
		t.Fatal(err)
	}
	checkGenerated(t, output, "zz_generated_selector_subscription_0.go", "SubscriptionFields")
	checkGenerated(t, output, "zz_generated_operations.go", "NewQuery", "NewSubscription")
}

func checkGenerated(t *testing.T, directory, file string, expected ...string) {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(directory, file))
	if err != nil {
		t.Fatal(err)
	}
	for _, text := range expected {
		if !strings.Contains(string(data), text) {
			t.Errorf("%s does not contain %q", file, text)
		}
	}
}

func generatedFiles(t *testing.T, directory string) []string {
	t.Helper()
	files, err := filepath.Glob(filepath.Join(directory, "*.go"))
	if err != nil {
		t.Fatal(err)
	}
	names := make([]string, len(files))
	for i, file := range files {
		names[i] = filepath.Base(file)
	}
	return names
}
