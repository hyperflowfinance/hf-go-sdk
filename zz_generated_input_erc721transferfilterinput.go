package hyperflow

type Erc721TransferFilterInput struct {
	TokenAddress   Optional[StringFilterInput]           `json:"tokenAddress"`
	FromAddress    Optional[StringFilterInput]           `json:"fromAddress"`
	ToAddress      Optional[StringFilterInput]           `json:"toAddress"`
	TokenID        Optional[StringFilterInput]           `json:"tokenId"`
	TxHash         Optional[StringFilterInput]           `json:"txHash"`
	BlockNumber    Optional[IntFilterInput]              `json:"blockNumber"`
	BlockTimestamp Optional[IntFilterInput]              `json:"blockTimestamp"`
	LogIndex       Optional[IntFilterInput]              `json:"logIndex"`
	ChainID        Optional[IntFilterInput]              `json:"chainId"`
	And            Optional[[]Erc721TransferFilterInput] `json:"and"`
	Or             Optional[[]Erc721TransferFilterInput] `json:"or"`
	Not            Optional[Erc721TransferFilterInput]   `json:"not"`
}

func (v Erc721TransferFilterInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
