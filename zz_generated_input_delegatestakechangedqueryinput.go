package hyperflow

type DelegateStakeChangedQueryInput struct {
	Filters    Optional[DelegateStakeChangedFilterInput]    `json:"filters"`
	Pagination Optional[PaginationInput]                    `json:"pagination"`
	OrderBy    Optional[[]DelegateStakeChangedOrderByInput] `json:"orderBy"`
}

func (v DelegateStakeChangedQueryInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
