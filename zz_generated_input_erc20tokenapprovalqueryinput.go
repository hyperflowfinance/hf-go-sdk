package hyperflow

type Erc20TokenApprovalQueryInput struct {
	Filters    Optional[Erc20TokenApprovalFilterInput]    `json:"filters"`
	Pagination Optional[PaginationInput]                  `json:"pagination"`
	OrderBy    Optional[[]Erc20TokenApprovalOrderByInput] `json:"orderBy"`
}

func (v Erc20TokenApprovalQueryInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
