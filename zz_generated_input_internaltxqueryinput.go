package hyperflow

type InternalTxQueryInput struct {
	Filters    Optional[InternalTxFilterInput]    `json:"filters"`
	Pagination Optional[PaginationInput]          `json:"pagination"`
	OrderBy    Optional[[]InternalTxOrderByInput] `json:"orderBy"`
}

func (v InternalTxQueryInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
