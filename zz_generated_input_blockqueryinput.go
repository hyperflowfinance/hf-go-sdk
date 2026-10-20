package hyperflow

type BlockQueryInput struct {
	Filters    Optional[BlockFilterInput]    `json:"filters"`
	Pagination Optional[PaginationInput]     `json:"pagination"`
	OrderBy    Optional[[]BlockOrderByInput] `json:"orderBy"`
}

func (v BlockQueryInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
