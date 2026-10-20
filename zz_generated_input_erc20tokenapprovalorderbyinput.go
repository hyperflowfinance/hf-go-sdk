package hyperflow

type Erc20TokenApprovalOrderByInput struct {
	Field     Erc20TokenApprovalOrderField `json:"field"`
	Direction SortDirection                `json:"direction"`
}

func (v Erc20TokenApprovalOrderByInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
