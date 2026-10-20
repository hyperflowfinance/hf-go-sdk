package hyperflow

type TransactionLeafInput struct {
	Raw       Optional[string]                `json:"raw"`
	TxData    Optional[BuildTransactionInput] `json:"txData"`
	Signature Optional[SignatureInput]        `json:"signature"`
}

func (v TransactionLeafInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
