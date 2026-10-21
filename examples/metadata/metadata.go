package metadata

import (
	"context"

	hyperflow "github.com/hyperflowfinance/hf-go-sdk"
)

func Get(ctx context.Context, client *hyperflow.Client) (hyperflow.Metadata, error) {
	query := hyperflow.NewQuery(func(fields *hyperflow.QueryFields) {
		fields.Metadata(func(metadata *hyperflow.MetadataFields) {
			metadata.ChainID()
		})
	})
	var result hyperflow.Query
	err := client.Query(ctx, query, &result)
	return result.Metadata, err
}
