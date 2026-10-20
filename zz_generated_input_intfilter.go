package hyperflow

type IntFilterInput struct {
	Eq      Optional[int32]   `json:"eq"`
	Ne      Optional[int32]   `json:"ne"`
	Gt      Optional[int32]   `json:"gt"`
	Gte     Optional[int32]   `json:"gte"`
	Lt      Optional[int32]   `json:"lt"`
	Lte     Optional[int32]   `json:"lte"`
	In      Optional[[]int32] `json:"in"`
	NotIn   Optional[[]int32] `json:"notIn"`
	Between Optional[[]int32] `json:"between"`
	IsNull  Optional[bool]    `json:"isNull"`
}

func (v IntFilterInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
