package hyperflow

type BlockOrderByInput struct {
	Field     BlockOrderField `json:"field"`
	Direction SortDirection   `json:"direction"`
}

func (v BlockOrderByInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
