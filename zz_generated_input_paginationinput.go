package hyperflow

type PaginationInput struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (v PaginationInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
