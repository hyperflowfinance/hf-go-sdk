package hyperflow

type Erc20TokenSwapEventOrderByInput struct {
	Field     Erc20TokenSwapEventOrderField `json:"field"`
	Direction SortDirection                 `json:"direction"`
}

func (v Erc20TokenSwapEventOrderByInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
