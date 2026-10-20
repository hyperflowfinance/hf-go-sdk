package hyperflow

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQueryBuildsTypedSelectionsAndOptionalVariables(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			Query     string         `json:"query"`
			Variables map[string]any `json:"variables"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Errorf("decode request: %v", err)
		}
		if request.Query != "query($v0: Int!) { search(limit: $v0) { ... on Thing { id } } }" {
			t.Errorf("unexpected query %q", request.Query)
		}
		if request.Variables["v0"] != float64(3) || len(request.Variables) != 1 {
			t.Errorf("unexpected variables %#v", request.Variables)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":{"search":[{"id":"t-1"}]}}`))
	}))
	defer server.Close()
	client, err := Connect(Config{URL: server.URL, HTTPClient: server.Client()})
	if err != nil {
		t.Fatal(err)
	}
	root := &selectionNode{}
	search := root.object("search",
		argument{name: "limit", gqlType: "Int!", value: int32(3)},
		argument{name: "filter", gqlType: "SearchFilter", value: Optional[json.RawMessage]{}, optional: true},
	)
	fragment := search.inlineFragment("Thing")
	fragment.field("id")
	query := Document{operation: "query", root: root}
	var result struct {
		Search []struct {
			ID string `json:"id"`
		} `json:"search"`
	}
	if err := client.Query(context.Background(), query, &result); err != nil {
		t.Fatal(err)
	}
	if len(result.Search) != 1 || result.Search[0].ID != "t-1" {
		t.Fatalf("unexpected result %#v", result)
	}
}
