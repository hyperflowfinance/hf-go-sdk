package hyperflow

type BigIntFilterInput struct {
	Eq      Optional[string]   `json:"eq"`
	Ne      Optional[string]   `json:"ne"`
	Gt      Optional[string]   `json:"gt"`
	Gte     Optional[string]   `json:"gte"`
	Lt      Optional[string]   `json:"lt"`
	Lte     Optional[string]   `json:"lte"`
	In      Optional[[]string] `json:"in"`
	NotIn   Optional[[]string] `json:"notIn"`
	Between Optional[[]string] `json:"between"`
	IsNull  Optional[bool]     `json:"isNull"`
}

func (v BigIntFilterInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
