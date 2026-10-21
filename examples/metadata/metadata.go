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
	if result.Metadata == nil {
		return hyperflow.Metadata{}, err
	}
	return *result.Metadata, err
}
