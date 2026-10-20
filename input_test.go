package hyperflow

import "testing"

type TestFilterFields struct {
	Status          Optional[string]   `json:"status"`
	Tags            Optional[[]string] `json:"tags"`
	IncludeArchived Optional[bool]     `json:"includeArchived"`
}

type TestFilterInput struct {
	TestFilterFields
}

func TestGeneratedInputTypesOmitUnsetFieldsAndKeepFalse(t *testing.T) {
	filter := TestFilterInput{TestFilterFields{
		Status: Some("ACTIVE"), Tags: Some([]string{"tag"}),
		IncludeArchived: Some(false),
	}}
	data, err := marshalInput(filter)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"includeArchived":false,"status":"ACTIVE","tags":["tag"]}`
	if string(data) != want {
		t.Fatalf("got %s, want %s", data, want)
	}
}
