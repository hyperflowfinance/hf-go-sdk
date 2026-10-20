package codegen

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func FetchSchema(ctx context.Context) (*Schema, error) {
	endpoint := os.Getenv("HYPERFLOW_API_URL")
	if endpoint == "" {
		endpoint = "https://api.hyperflowlabs.com/graphql"
	}
	body, err := json.Marshal(map[string]string{"query": introspectionQuery()})
	if err != nil {
		return nil, fmt.Errorf("encode schema request: %w", err)
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create schema request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")
	auth := os.Getenv("HYPERFLOW_AUTH_VALUE")
	if auth != "" {
		header := os.Getenv("HYPERFLOW_AUTH_HEADER")
		if header == "" {
			header = "x-api-key"
		}
		request.Header.Set(header, auth)
	}
	response, err := (&http.Client{Timeout: 30 * time.Second}).Do(request)
	if err != nil {
		return nil, fmt.Errorf("fetch GraphQL schema: %w", err)
	}
	defer response.Body.Close()
	data, err := io.ReadAll(io.LimitReader(response.Body, (16<<20)+1))
	if err != nil {
		return nil, fmt.Errorf("read schema response: %w", err)
	}
	if len(data) > 16<<20 {
		return nil, fmt.Errorf("schema introspection response exceeds 16 MiB")
	}
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, fmt.Errorf("schema introspection failed: HTTP %d: %s", response.StatusCode, data)
	}
	return ParseSchema(bytes.NewReader(data))
}
