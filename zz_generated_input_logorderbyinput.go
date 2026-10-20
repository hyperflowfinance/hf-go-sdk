package hyperflow

type LogOrderByInput struct {
	Field     LogOrderField `json:"field"`
	Direction SortDirection `json:"direction"`
}

func (v LogOrderByInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
