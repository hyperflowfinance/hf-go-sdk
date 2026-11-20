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

See [examples/metadata-cli](examples/metadata-cli) for queries and
[examples/new-blocks](examples/new-blocks) for WebSocket subscriptions. Decode into generated `Query`/`SubscriptionEvent` structs or caller-provided structs; only selected fields are populated.
Selectors, arguments, enums, inputs and response types come from the live schema. Nullable outputs use pointers, lists use slices, and custom scalars use `json.RawMessage` (`metadata.chainId` is a nullable hexadecimal string).
`NewSubscription` builds typed documents; `Client.Subscribe(ctx, document)` returns a stream. Call `Next(&event)` until `io.EOF`, then `Close()`. Subscriptions use `graphql-transport-ws`; the example prints JSON events with nullable fields as `null`.

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

`Connect` accepts a `Config`; explicit values override environment variables. `URL` defaults to production, `AuthHeader` to `x-api-key`, and `AuthValue` to `HYPERFLOW_AUTH_VALUE`.
Environment names are `HYPERFLOW_API_URL`, `HYPERFLOW_AUTH_HEADER` and `HYPERFLOW_AUTH_VALUE`. Pass an empty `AuthValue` pointer to disable environment auth; custom `HTTPClient` and `Headers` are preserved.

## Generate and test

```sh
go run ./cmd/codegen
go test ./...
go run ./examples/metadata-cli
```

Generation introspects live and fails on HTTP 429 or GraphQL errors; release builds can use a verified stage snapshot. Generated files and downloaded schemas are ignored on `main`; CI generates once before testing.
Selector files split at the 99-line limit; a single oversized method fails explicitly.

## Go module releases

Go has no upload registry. Dispatch `.github/workflows/release.yml` from `main` with an approved stable `vX.Y.Z`; it tests generated code, creates a release-only commit, pushes the tag, then requests indexing from `proxy.golang.org`.
No registry token is used. Do not edit or retag an existing version.

Stage changes publish a deterministic prerelease tag from the verified snapshot. Use the version in the manifest SDK status with `go get github.com/hyperflowfinance/hf-go-sdk@<version>`.
Preview tags default to stage; stable versions remain production-default.

Verify the version and schema hash against the manifest, run `go test ./...`,
then smoke-test a read-only query with `go run ./examples/metadata-cli` against
`https://stage.api.hyperflowlabs.com/graphql`. Promote only after the stage query
succeeds; stable tags remain production-default.
