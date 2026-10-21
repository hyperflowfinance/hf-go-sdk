package hyperflow

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

const defaultURL = "https://api.hyperflowlabs.com/graphql"

// Config configures the GraphQL endpoint, authentication, headers, and HTTP client.
type Config struct {
	URL        string
	AuthHeader string
	AuthValue  *string
	Headers    http.Header
	HTTPClient *http.Client
}

// Client sends typed GraphQL queries, mutations, and subscriptions.
type Client struct {
	endpoint string
	headers  http.Header
	http     *http.Client
}

// Connect creates a GraphQL client using the supplied configuration and environment defaults.
func Connect(config ...Config) (*Client, error) {
	cfg := Config{}
	if len(config) > 1 {
		return nil, fmt.Errorf("Connect accepts at most one Config")
	}
	if len(config) == 1 {
		cfg = config[0]
	}
	endpoint := cfg.URL
	if endpoint == "" {
		endpoint = os.Getenv("HYPERFLOW_API_URL")
	}
	if endpoint == "" {
		endpoint = defaultURL
	}
	parsed, err := url.ParseRequestURI(endpoint)
	if err != nil || parsed.Host == "" || (parsed.Scheme != "http" && parsed.Scheme != "https") {
		return nil, fmt.Errorf("invalid GraphQL endpoint %q", endpoint)
	}
	header := cfg.AuthHeader
	if header == "" {
		header = os.Getenv("HYPERFLOW_AUTH_HEADER")
	}
	if header == "" {
		header = "x-api-key"
	}
	auth := os.Getenv("HYPERFLOW_AUTH_VALUE")
	if cfg.AuthValue != nil {
		auth = *cfg.AuthValue
	}
	headers := cfg.Headers.Clone()
	if headers == nil {
		headers = make(http.Header)
	}
	if auth != "" {
		headers.Set(header, auth)
	}
	client := cfg.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{endpoint: endpoint, headers: headers, http: client}, nil
}
