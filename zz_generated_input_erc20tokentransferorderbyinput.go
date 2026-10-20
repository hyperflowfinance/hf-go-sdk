package hyperflow

type Erc20TokenTransferOrderByInput struct {
	Field     Erc20TokenTransferOrderField `json:"field"`
	Direction SortDirection                `json:"direction"`
}

func (v Erc20TokenTransferOrderByInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
