// Package hyperflow builds schema-typed GraphQL operations for the Hyperflow API.
//
// Generate the schema selectors with go run ./cmd/codegen, then build an operation
// with NewQuery, NewMutation, or NewSubscription. Client.Query and Client.Mutate
// decode the selected response fields into a caller-defined Go struct; this keeps
// response types explicit and lets each struct match the requested selection.
//
// For example, define a result type with JSON tags:
//
//	type MetadataResult struct {
//		Metadata struct {
//			ChainID string `json:"chainId"`
//		} `json:"metadata"`
//	}
//
// Use hyperflow.Some to provide optional GraphQL arguments, or the zero value of
// Optional[T] to omit them. Connect reads endpoint and authentication defaults
// from HYPERFLOW_API_URL, HYPERFLOW_AUTH_HEADER, and HYPERFLOW_AUTH_VALUE.
//
// For a runnable query, see ExampleClient_Query. Subscriptions use
// Client.Subscribe and Subscription.Next over the graphql-transport-ws protocol.
package hyperflow
