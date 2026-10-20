package codegen

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFetchReportsRateLimitWithoutRetry(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		calls++
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte("rate limited"))
	}))
	defer server.Close()
	t.Setenv("HYPERFLOW_API_URL", server.URL)
	_, err := FetchSchema(context.Background())
	if err == nil || !strings.Contains(err.Error(), "HTTP 429") {
		t.Fatalf("expected explicit HTTP 429, got %v", err)
	}
	if calls != 1 {
		t.Fatalf("made %d requests, want one", calls)
	}
}

func TestParseSchemaReportsGraphQLErrors(t *testing.T) {
	_, err := ParseSchema(strings.NewReader(`{"errors":[{"message":"denied"}]}`))
	if err == nil || !strings.Contains(err.Error(), "denied") {
		t.Fatalf("expected GraphQL error, got %v", err)
	}
}
