package metadata

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	hyperflow "github.com/hyperflowfinance/hf-go-sdk"
)

func TestExampleUsesTypedSelectionAndConfiguredClient(t *testing.T) {
	t.Setenv("HYPERFLOW_API_URL", "")
	t.Setenv("HYPERFLOW_AUTH_HEADER", "")
	t.Setenv("HYPERFLOW_AUTH_VALUE", "")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Test") != "kept" {
			t.Errorf("custom header was lost")
		}
		var request struct {
			Query string `json:"query"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Errorf("decode request: %v", err)
		}
		if request.Query != "query { metadata { chainId } }" {
			t.Errorf("unexpected query %q", request.Query)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":{"metadata":{"chainId":"1"}}}`))
	}))
	defer server.Close()
	client, err := hyperflow.Connect(hyperflow.Config{
		URL: server.URL, Headers: http.Header{"X-Test": []string{"kept"}},
		HTTPClient: server.Client(),
	})
	if err != nil {
		t.Fatal(err)
	}
	result, err := Get(context.Background(), client)
	if err != nil {
		t.Fatal(err)
	}
	if result.ChainID != "1" {
		t.Fatalf("unexpected chain ID %q", result.ChainID)
	}
}
