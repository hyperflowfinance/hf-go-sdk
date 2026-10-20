package hyperflow

type LogQueryInput struct {
	Filters    Optional[LogFilterInput]    `json:"filters"`
	Pagination Optional[PaginationInput]   `json:"pagination"`
	OrderBy    Optional[[]LogOrderByInput] `json:"orderBy"`
}

func (v LogQueryInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
