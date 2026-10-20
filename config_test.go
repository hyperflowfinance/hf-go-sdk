package hyperflow

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConfigOverridesEnvironmentAndCanDisableAuth(t *testing.T) {
	t.Setenv("HYPERFLOW_API_URL", "https://env.invalid/graphql")
	t.Setenv("HYPERFLOW_AUTH_HEADER", "X-Env-Key")
	t.Setenv("HYPERFLOW_AUTH_VALUE", "env-secret")
	for _, test := range []struct {
		name, auth string
		want       string
	}{
		{name: "explicit auth", auth: "explicit-secret", want: "explicit-secret"},
		{name: "explicit empty auth", auth: "", want: ""},
	} {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if got := r.Header.Get("X-Explicit"); got != test.want {
					t.Errorf("auth header = %q, want %q", got, test.want)
				}
				if got := r.Header.Get("X-Env-Key"); got != "" {
					t.Errorf("environment auth leaked through override: %q", got)
				}
				_, _ = w.Write([]byte(`{"data":{"metadata":{"chainId":"1"}}}`))
			}))
			defer server.Close()
			auth := test.auth
			client, err := Connect(Config{
				URL: server.URL, AuthHeader: "X-Explicit", AuthValue: &auth,
				HTTPClient: server.Client(),
			})
			if err != nil {
				t.Fatal(err)
			}
			if err := client.Query(context.Background(), testQueryDocument(), &struct{}{}); err != nil {
				t.Fatal(err)
			}
		})
	}
}
