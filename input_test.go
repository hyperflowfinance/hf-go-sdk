package hyperflow

import (
	"encoding/json"
	"testing"
)

type TestFilterFields struct {
	Status          Optional[string]   `json:"status,omitzero"`
	Tags            Optional[[]string] `json:"tags,omitzero"`
	IncludeArchived Optional[bool]     `json:"includeArchived,omitzero"`
}

type TestFilterInput struct {
	TestFilterFields
}

func TestGeneratedInputTypesOmitUnsetFieldsAndKeepFalse(t *testing.T) {
	filter := TestFilterInput{TestFilterFields{
		Status: Some("ACTIVE"), Tags: Some([]string{"tag"}),
		IncludeArchived: Some(false),
	}}
	data, err := json.Marshal(filter)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"status":"ACTIVE","tags":["tag"],"includeArchived":false}`
	if string(data) != want {
		t.Fatalf("got %s, want %s", data, want)
	}
}

func TestOptionalSupportsRecursiveInputTypes(t *testing.T) {
	type recursiveInput struct {
		Next Optional[recursiveInput] `json:"next,omitzero"`
	}
	data, err := json.Marshal(recursiveInput{Next: Some(recursiveInput{})})
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != `{"next":{}}` {
		t.Fatalf("got %s, want recursive input object", data)
	}
}
