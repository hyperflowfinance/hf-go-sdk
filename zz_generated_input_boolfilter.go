package hyperflow

type BoolFilterInput struct {
	Eq     Optional[bool] `json:"eq"`
	IsNull Optional[bool] `json:"isNull"`
}

func (v BoolFilterInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
