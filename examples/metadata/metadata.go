package metadata

import (
	"context"

	hyperflow "github.com/hyperflowfinance/hf-go-sdk"
)

type Result struct {
	Metadata struct {
		ChainID string `json:"chainId"`
	} `json:"metadata"`
}

func Get(ctx context.Context, client *hyperflow.Client) (Result, error) {
	query := hyperflow.NewQuery(func(fields *hyperflow.QueryFields) {
		fields.Metadata(func(metadata *hyperflow.MetadataFields) {
			metadata.ChainID(hyperflow.Optional[string]{}, hyperflow.Optional[int32]{})
		})
	})
	var result Result
	err := client.Query(ctx, query, &result)
	return result, err
}
