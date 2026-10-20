package hyperflow

type Erc20TokenTransferQueryInput struct {
	Filters    Optional[Erc20TokenTransferFilterInput]    `json:"filters"`
	Pagination Optional[PaginationInput]                  `json:"pagination"`
	OrderBy    Optional[[]Erc20TokenTransferOrderByInput] `json:"orderBy"`
}

func (v Erc20TokenTransferQueryInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
