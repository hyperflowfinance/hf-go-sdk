package hyperflow

type TransactionQueryInput struct {
	Filters    Optional[TransactionFilterInput]    `json:"filters"`
	Pagination Optional[PaginationInput]           `json:"pagination"`
	OrderBy    Optional[[]TransactionOrderByInput] `json:"orderBy"`
}

func (v TransactionQueryInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
