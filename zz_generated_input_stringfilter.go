package hyperflow

type StringFilterInput struct {
	Eq          Optional[string]   `json:"eq"`
	Ne          Optional[string]   `json:"ne"`
	Contains    Optional[string]   `json:"contains"`
	NotContains Optional[string]   `json:"notContains"`
	StartsWith  Optional[string]   `json:"startsWith"`
	EndsWith    Optional[string]   `json:"endsWith"`
	Regex       Optional[string]   `json:"regex"`
	NotRegex    Optional[string]   `json:"notRegex"`
	In          Optional[[]string] `json:"in"`
	NotIn       Optional[[]string] `json:"notIn"`
	IsNull      Optional[bool]     `json:"isNull"`
}

func (v StringFilterInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
