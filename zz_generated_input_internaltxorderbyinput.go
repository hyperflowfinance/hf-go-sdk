package hyperflow

type InternalTxOrderByInput struct {
	Field     InternalTxOrderField `json:"field"`
	Direction SortDirection        `json:"direction"`
}

func (v InternalTxOrderByInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
