package hyperflow

type DelegatedOrderByInput struct {
	Field     DelegatedOrderField `json:"field"`
	Direction SortDirection       `json:"direction"`
}

func (v DelegatedOrderByInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
