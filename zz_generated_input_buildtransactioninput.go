package hyperflow

type BuildTransactionInput struct {
	To                   string           `json:"to"`
	Value                string           `json:"value"`
	Gas                  string           `json:"gas"`
	Nonce                string           `json:"nonce"`
	Data                 string           `json:"data"`
	ChainID              string           `json:"chainId"`
	GasPrice             Optional[string] `json:"gasPrice"`
	MaxFeePerGas         Optional[string] `json:"maxFeePerGas"`
	MaxPriorityFeePerGas Optional[string] `json:"maxPriorityFeePerGas"`
	Type                 Optional[string] `json:"type"`
}

func (v BuildTransactionInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
