package hyperflow

type InternalTxFilterInput struct {
	BlockNumber  Optional[IntFilterInput]          `json:"blockNumber"`
	TxIndex      Optional[IntFilterInput]          `json:"txIndex"`
	TraceIndex   Optional[IntFilterInput]          `json:"traceIndex"`
	Subtraces    Optional[IntFilterInput]          `json:"subtraces"`
	TxHash       Optional[StringFilterInput]       `json:"txHash"`
	TraceAddress Optional[StringFilterInput]       `json:"traceAddress"`
	FromAddress  Optional[StringFilterInput]       `json:"fromAddress"`
	ToAddress    Optional[StringFilterInput]       `json:"toAddress"`
	CallType     Optional[StringFilterInput]       `json:"callType"`
	Input        Optional[StringFilterInput]       `json:"input"`
	Error        Optional[StringFilterInput]       `json:"error"`
	ValueWei     Optional[BigIntFilterInput]       `json:"valueWei"`
	Gas          Optional[BigIntFilterInput]       `json:"gas"`
	GasUsed      Optional[BigIntFilterInput]       `json:"gasUsed"`
	Success      Optional[IntFilterInput]          `json:"success"`
	ChainID      Optional[IntFilterInput]          `json:"chainId"`
	And          Optional[[]InternalTxFilterInput] `json:"and"`
	Or           Optional[[]InternalTxFilterInput] `json:"or"`
	Not          Optional[InternalTxFilterInput]   `json:"not"`
}

func (v InternalTxFilterInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
