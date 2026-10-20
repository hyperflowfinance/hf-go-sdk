package hyperflow

type Erc20TokenApprovalFilterInput struct {
	TokenAddress   Optional[StringFilterInput]               `json:"tokenAddress"`
	Owner          Optional[StringFilterInput]               `json:"owner"`
	Spender        Optional[StringFilterInput]               `json:"spender"`
	Value          Optional[StringFilterInput]               `json:"value"`
	TxHash         Optional[StringFilterInput]               `json:"txHash"`
	BlockNumber    Optional[IntFilterInput]                  `json:"blockNumber"`
	BlockTimestamp Optional[IntFilterInput]                  `json:"blockTimestamp"`
	LogIndex       Optional[IntFilterInput]                  `json:"logIndex"`
	ChainID        Optional[IntFilterInput]                  `json:"chainId"`
	And            Optional[[]Erc20TokenApprovalFilterInput] `json:"and"`
	Or             Optional[[]Erc20TokenApprovalFilterInput] `json:"or"`
	Not            Optional[Erc20TokenApprovalFilterInput]   `json:"not"`
}

func (v Erc20TokenApprovalFilterInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
