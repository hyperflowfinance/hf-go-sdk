package hyperflow

type EstimateGasInput struct {
	To                   Optional[string] `json:"to"`
	FromAddress          Optional[string] `json:"fromAddress"`
	Gas                  Optional[string] `json:"gas"`
	GasPrice             Optional[string] `json:"gasPrice"`
	MaxFeePerGas         Optional[string] `json:"maxFeePerGas"`
	MaxPriorityFeePerGas Optional[string] `json:"maxPriorityFeePerGas"`
	Value                Optional[string] `json:"value"`
	Data                 Optional[string] `json:"data"`
	BlockIdentifier      Optional[string] `json:"blockIdentifier"`
}

func (v EstimateGasInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
