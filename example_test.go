package hyperflow_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	hyperflow "github.com/hyperflowfinance/hf-go-sdk"
)

type MetadataResult struct {
	Metadata struct {
		ChainID string `json:"chainId"`
	} `json:"metadata"`
}

func ExampleClient_Query() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `{"data":{"metadata":{"chainId":"1"}}}`)
	}))
	defer server.Close()

	noAuth := ""
	client, err := hyperflow.Connect(hyperflow.Config{URL: server.URL, AuthValue: &noAuth})
	if err != nil {
		panic(err)
	}
	query := hyperflow.NewQuery(func(q *hyperflow.QueryFields) {
		q.Metadata(func(m *hyperflow.MetadataFields) {
			m.ChainID(hyperflow.Optional[string]{}, hyperflow.Optional[int32]{})
		})
	})
	var result MetadataResult
	if err := client.Query(context.Background(), query, &result); err != nil {
		panic(err)
	}
	fmt.Println(result.Metadata.ChainID)
	// Output: 1
}
