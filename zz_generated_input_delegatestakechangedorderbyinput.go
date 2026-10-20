package hyperflow

type DelegateStakeChangedOrderByInput struct {
	Field     DelegateStakeChangedOrderField `json:"field"`
	Direction SortDirection                  `json:"direction"`
}

func (v DelegateStakeChangedOrderByInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
