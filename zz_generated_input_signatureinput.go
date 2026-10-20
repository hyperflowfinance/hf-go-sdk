package hyperflow

type SignatureInput struct {
	R string `json:"r"`
	S string `json:"s"`
	V string `json:"v"`
}

func (v SignatureInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
