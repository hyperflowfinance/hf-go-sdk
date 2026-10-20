package hyperflow

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHTTPAndGraphQLErrorsAreReturned(t *testing.T) {
	for _, test := range []struct {
		status int
		body   string
		want   string
	}{
		{http.StatusTooManyRequests, "limited", "GraphQL HTTP 429"},
		{http.StatusOK, `{"errors":[{"message":"denied"}]}`, "GraphQL errors: denied"},
	} {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(test.status)
			_, _ = w.Write([]byte(test.body))
		}))
		client, err := Connect(Config{URL: server.URL, HTTPClient: server.Client()})
		if err != nil {
			t.Fatal(err)
		}
		err = client.Query(context.Background(), testQueryDocument(), &struct{}{})
		server.Close()
		if err == nil || !strings.Contains(err.Error(), test.want) {
			t.Errorf("expected %q, got %v", test.want, err)
		}
		if test.status == http.StatusTooManyRequests {
			var httpErr *HTTPError
			if !errors.As(err, &httpErr) {
				t.Errorf("expected HTTPError, got %T", err)
			}
		}
	}
}

func TestClientRejectsInvalidResultAndOperation(t *testing.T) {
	client, err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	query := Document{operation: "mutation", root: &selectionNode{}}
	if err := client.Query(context.Background(), query, &struct{}{}); err == nil {
		t.Fatal("expected operation mismatch")
	}
	var result *struct{}
	if err := client.Query(context.Background(), Document{operation: "query"}, result); err == nil {
		t.Fatal("expected nil result error")
	}
}
