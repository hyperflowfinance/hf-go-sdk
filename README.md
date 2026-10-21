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
        m.ChainID()
    })
})
var result hyperflow.Query
if err := client.Query(ctx, query, &result); err != nil {
    return err
}
if result.Metadata.ChainID == nil {
    return fmt.Errorf("metadata.chainId is null")
}
fmt.Println(*result.Metadata.ChainID)
```

See [examples/metadata-cli](examples/metadata-cli) for a query example and
[examples/new-blocks](examples/new-blocks) for a WebSocket subscription example.
Results can be decoded into generated response structs such as `Query` and
`SubscriptionEvent`, or caller-provided structs matching the selected response
shape. Generated response structs contain all schema fields, but only selected
fields are populated; unselected fields keep their Go zero values.
Generated field selectors, argument types, enums, input types and response types
come from the live schema. Nullable scalar and object outputs use pointers;
lists use slices, and custom scalars use `json.RawMessage`. For example,
`metadata.chainId` is a nullable hexadecimal string, not an integer.
When available, `NewSubscription` builds a schema-typed subscription document.
`Client.Subscribe(ctx, document)` returns a stream; call `Next(&event)` for each
result until it returns `io.EOF`, then call `Close()`. Subscriptions use GraphQL's
`graphql-transport-ws` protocol over WebSocket. The new-blocks example prints one
JSON object per event, preserving nullable fields as `null`.

## Field selection and arguments

Selection callbacks choose nested GraphQL fields; they do not process results
or send requests. Required arguments are positional. Optional arguments go in
the field's generated `Args` struct after the selection callback:

```go
query := hyperflow.NewQuery(func(q *hyperflow.QueryFields) {
    q.Blocks(func(connection *hyperflow.BlockConnectionFields) {
        connection.Items(func(b *hyperflow.BlockFields) {
            b.Number()
            b.Hash()
        })
    }, hyperflow.QueryBlocksArgs{
        QueryInput: hyperflow.Some(hyperflow.BlockQueryInput{
            Pagination: hyperflow.Some(hyperflow.PaginationInput{Limit: 5}),
        }),
    })
})
```

Omit the `Args` value to use server defaults. Nullable inputs use `Optional[T]`:
the zero value omits an input, `Some(value)` supplies it (including `false` or
`0`), and `Optional[T]{Set: true}` supplies explicit `null`.

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
Selector files are split by formatted line count to stay within the 99-line
generated-file limit. A single oversized method still fails explicitly.

## Go module releases

Go has no npm-style upload registry. Dispatch `.github/workflows/release.yml`
from `main` with an approved `vX.Y.Z` version. It tests the generated client,
creates a release-only commit containing generated sources, pushes that tag,
and asks `proxy.golang.org` to index the module. No registry token or upload
step is used. Do not edit or retag an existing version.
