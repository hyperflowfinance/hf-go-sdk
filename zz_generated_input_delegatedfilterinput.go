package hyperflow

type DelegatedFilterInput struct {
	ContractAddress Optional[StringFilterInput]      `json:"contractAddress"`
	Delegator       Optional[StringFilterInput]      `json:"delegator"`
	Guardian        Optional[StringFilterInput]      `json:"guardian"`
	TxHash          Optional[StringFilterInput]      `json:"txHash"`
	BlockNumber     Optional[IntFilterInput]         `json:"blockNumber"`
	BlockTimestamp  Optional[IntFilterInput]         `json:"blockTimestamp"`
	LogIndex        Optional[IntFilterInput]         `json:"logIndex"`
	ChainID         Optional[IntFilterInput]         `json:"chainId"`
	And             Optional[[]DelegatedFilterInput] `json:"and"`
	Or              Optional[[]DelegatedFilterInput] `json:"or"`
	Not             Optional[DelegatedFilterInput]   `json:"not"`
}

func (v DelegatedFilterInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
