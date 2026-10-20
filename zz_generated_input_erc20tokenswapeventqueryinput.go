package hyperflow

type Erc20TokenSwapEventQueryInput struct {
	Filters    Optional[Erc20TokenSwapEventFilterInput]    `json:"filters"`
	Pagination Optional[PaginationInput]                   `json:"pagination"`
	OrderBy    Optional[[]Erc20TokenSwapEventOrderByInput] `json:"orderBy"`
}

func (v Erc20TokenSwapEventQueryInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
