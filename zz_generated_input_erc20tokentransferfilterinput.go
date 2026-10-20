package hyperflow

type Erc20TokenTransferFilterInput struct {
	TokenAddress   Optional[StringFilterInput]               `json:"tokenAddress"`
	FromAddress    Optional[StringFilterInput]               `json:"fromAddress"`
	ToAddress      Optional[StringFilterInput]               `json:"toAddress"`
	Value          Optional[StringFilterInput]               `json:"value"`
	TxHash         Optional[StringFilterInput]               `json:"txHash"`
	BlockNumber    Optional[IntFilterInput]                  `json:"blockNumber"`
	BlockTimestamp Optional[IntFilterInput]                  `json:"blockTimestamp"`
	LogIndex       Optional[IntFilterInput]                  `json:"logIndex"`
	ChainID        Optional[IntFilterInput]                  `json:"chainId"`
	And            Optional[[]Erc20TokenTransferFilterInput] `json:"and"`
	Or             Optional[[]Erc20TokenTransferFilterInput] `json:"or"`
	Not            Optional[Erc20TokenTransferFilterInput]   `json:"not"`
}

func (v Erc20TokenTransferFilterInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
