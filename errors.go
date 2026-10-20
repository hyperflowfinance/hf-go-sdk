package hyperflow

import (
	"fmt"
	"strings"
)

type HTTPError struct {
	StatusCode int
	Body       string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("GraphQL HTTP %d: %s", e.StatusCode, e.Body)
}

type GraphQLError struct {
	Message    string         `json:"message"`
	Path       []any          `json:"path,omitempty"`
	Extensions map[string]any `json:"extensions,omitempty"`
}

type GraphQLErrors []GraphQLError

func (e GraphQLErrors) Error() string {
	messages := make([]string, 0, len(e))
	for _, item := range e {
		messages = append(messages, item.Message)
	}
	return "GraphQL errors: " + strings.Join(messages, "; ")
}
