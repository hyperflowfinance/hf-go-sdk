package hyperflow

type GetTransactionCountInput struct {
	Address         string          `json:"address"`
	BlockIdentifier string          `json:"blockIdentifier"`
	ChainID         Optional[int32] `json:"chainId"`
}

func (v GetTransactionCountInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
