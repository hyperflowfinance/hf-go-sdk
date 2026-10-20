package hyperflow

type DelegateStakeChangedFilterInput struct {
	ContractAddress Optional[StringFilterInput]                 `json:"contractAddress"`
	Delegator       Optional[StringFilterInput]                 `json:"delegator"`
	Guardian        Optional[StringFilterInput]                 `json:"guardian"`
	TxHash          Optional[StringFilterInput]                 `json:"txHash"`
	BlockNumber     Optional[IntFilterInput]                    `json:"blockNumber"`
	BlockTimestamp  Optional[IntFilterInput]                    `json:"blockTimestamp"`
	LogIndex        Optional[IntFilterInput]                    `json:"logIndex"`
	ChainID         Optional[IntFilterInput]                    `json:"chainId"`
	And             Optional[[]DelegateStakeChangedFilterInput] `json:"and"`
	Or              Optional[[]DelegateStakeChangedFilterInput] `json:"or"`
	Not             Optional[DelegateStakeChangedFilterInput]   `json:"not"`
}

func (v DelegateStakeChangedFilterInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
