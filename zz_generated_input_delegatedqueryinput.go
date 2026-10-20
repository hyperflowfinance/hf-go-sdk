package hyperflow

type DelegatedQueryInput struct {
	Filters    Optional[DelegatedFilterInput]    `json:"filters"`
	Pagination Optional[PaginationInput]         `json:"pagination"`
	OrderBy    Optional[[]DelegatedOrderByInput] `json:"orderBy"`
}

func (v DelegatedQueryInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
