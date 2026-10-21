package hyperflow

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

// Query executes a generated query document and decodes data into result.
func (c *Client) Query(ctx context.Context, document Document, result any) error {
	return c.execute(ctx, "query", document, result)
}

// Mutate executes a generated mutation document and decodes data into result.
func (c *Client) Mutate(ctx context.Context, document Document, result any) error {
	return c.execute(ctx, "mutation", document, result)
}

func (c *Client) execute(ctx context.Context, operation string, document Document, result any) error {
	if document.operation != operation {
		return fmt.Errorf("expected %s document", operation)
	}
	value := reflect.ValueOf(result)
	if !value.IsValid() || value.Kind() != reflect.Pointer || value.IsNil() {
		return fmt.Errorf("result must be a non-nil pointer")
	}
	query, variables, err := document.render()
	if err != nil {
		return err
	}
	body, err := json.Marshal(map[string]any{"query": query, "variables": normalize(reflect.ValueOf(variables))})
	if err != nil {
		return fmt.Errorf("encode GraphQL request: %w", err)
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("create GraphQL request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")
	for key, values := range c.headers {
		for _, value := range values {
			request.Header.Add(key, value)
		}
	}
	return c.decodeResponse(request, result)
}

func (c *Client) decodeResponse(request *http.Request, result any) error {
	response, err := c.http.Do(request)
	if err != nil {
		return fmt.Errorf("send GraphQL request: %w", err)
	}
	defer response.Body.Close()
	data, err := io.ReadAll(io.LimitReader(response.Body, (4<<20)+1))
	if err != nil {
		return fmt.Errorf("read GraphQL response: %w", err)
	}
	if len(data) > 4<<20 {
		return fmt.Errorf("GraphQL response exceeds 4 MiB")
	}
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return &HTTPError{StatusCode: response.StatusCode, Body: string(data)}
	}
	var decoded struct {
		Data   json.RawMessage `json:"data"`
		Errors GraphQLErrors   `json:"errors"`
	}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return fmt.Errorf("decode GraphQL response: %w", err)
	}
	if len(decoded.Errors) != 0 {
		return decoded.Errors
	}
	if len(decoded.Data) == 0 || string(decoded.Data) == "null" {
		return fmt.Errorf("GraphQL response is missing data")
	}
	if err := json.Unmarshal(decoded.Data, result); err != nil {
		return fmt.Errorf("decode GraphQL data: %w", err)
	}
	return nil
}
