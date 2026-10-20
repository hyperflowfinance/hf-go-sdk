package hyperflow

type TransactionNodeInput struct {
	Item Optional[TransactionLeafInput]   `json:"item"`
	Seq  Optional[[]TransactionNodeInput] `json:"seq"`
	Par  Optional[[]TransactionNodeInput] `json:"par"`
}

func (v TransactionNodeInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
