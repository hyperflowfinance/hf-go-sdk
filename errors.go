package hyperflow

import (
	"fmt"
	"strings"
)

// HTTPError reports a non-success HTTP response from the GraphQL endpoint.
type HTTPError struct {
	StatusCode int
	Body       string
}

// Error formats the failed GraphQL HTTP response.
func (e *HTTPError) Error() string {
	return fmt.Sprintf("GraphQL HTTP %d: %s", e.StatusCode, e.Body)
}

// GraphQLError is an error returned by the GraphQL server.
type GraphQLError struct {
	Message    string         `json:"message"`
	Path       []any          `json:"path,omitempty"`
	Extensions map[string]any `json:"extensions,omitempty"`
}

// GraphQLErrors contains GraphQL errors returned by the server.
type GraphQLErrors []GraphQLError

// Error joins the messages in the GraphQL error response.
func (e GraphQLErrors) Error() string {
	messages := make([]string, 0, len(e))
	for _, item := range e {
		messages = append(messages, item.Message)
	}
	return "GraphQL errors: " + strings.Join(messages, "; ")
}
