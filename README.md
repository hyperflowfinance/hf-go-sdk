# hf-go-sdk

Typed Go client and schema-wide query builder for the HyperFlow GraphQL API.
Requires Go 1.26 or newer.

## Use

```go
client, err := hyperflow.Connect()
if err != nil {
    return err
}
query := hyperflow.NewQuery(func(q *hyperflow.QueryFields) {
    q.Metadata(func(m *hyperflow.MetadataFields) {
        m.ChainID(hyperflow.Optional[string]{}, hyperflow.Optional[int32]{})
    })
})
var result struct {
    Metadata struct { ChainID string `json:"chainId"` } `json:"metadata"`
}
if err := client.Query(ctx, query, &result); err != nil {
    return err
}
fmt.Println(result.Metadata.ChainID)
```

See [examples/metadata-cli](examples/metadata-cli) for a query example and
[examples/new-blocks](examples/new-blocks) for a WebSocket subscription example.
Results are decoded into caller-provided structs, matching the selected response
shape.
Generated field selectors, argument types, enums and input types come from the
live schema. Custom scalars use `json.RawMessage`; nullable inputs use `Optional[T]`.
When available, `NewSubscription` builds a schema-typed subscription document.
`Client.Subscribe(ctx, document)` returns a stream; call `Next(&event)` for each
result until it returns `io.EOF`, then call `Close()`. Subscriptions use GraphQL's
`graphql-transport-ws` protocol over WebSocket.

## Configuration

`Connect` accepts a `Config`; explicit values override environment variables.
`URL` defaults to `https://api.hyperflowlabs.com/graphql`. `AuthHeader` defaults
to `x-api-key`; `AuthValue` defaults to `HYPERFLOW_AUTH_VALUE`. The environment
names are `HYPERFLOW_API_URL`, `HYPERFLOW_AUTH_HEADER` and `HYPERFLOW_AUTH_VALUE`.
Pass a pointer to an empty `AuthValue` to disable environment authentication.
Custom `HTTPClient` and `Headers` are preserved.

## Generate and test

```sh
go run ./cmd/codegen
go test ./...
go run ./examples/metadata-cli
```

Generation introspects the live endpoint and fails explicitly on HTTP 429 or
GraphQL errors; it does not retry. Generated `zz_generated_*.go` files and the
downloaded schema are ignored on `main`. CI generates once before testing.

## Go module releases

Go has no npm-style upload registry. Dispatch `.github/workflows/release.yml`
from `main` with an approved `vX.Y.Z` version. It tests the generated client,
creates a release-only commit containing generated sources, pushes that tag,
and asks `proxy.golang.org` to index the module. No registry token or upload
step is used. Do not edit or retag an existing version.
