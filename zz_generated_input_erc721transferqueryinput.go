package hyperflow

type Erc721TransferQueryInput struct {
	Filters    Optional[Erc721TransferFilterInput]    `json:"filters"`
	Pagination Optional[PaginationInput]              `json:"pagination"`
	OrderBy    Optional[[]Erc721TransferOrderByInput] `json:"orderBy"`
}

func (v Erc721TransferQueryInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
