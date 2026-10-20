package hyperflow

type TransactionOrderByInput struct {
	Field     TransactionOrderField `json:"field"`
	Direction SortDirection         `json:"direction"`
}

func (v TransactionOrderByInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
