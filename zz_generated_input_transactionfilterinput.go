package hyperflow

type TransactionFilterInput struct {
	Hash                 Optional[StringFilterInput]        `json:"hash"`
	FromAddress          Optional[StringFilterInput]        `json:"fromAddress"`
	ToAddress            Optional[StringFilterInput]        `json:"toAddress"`
	Input                Optional[StringFilterInput]        `json:"input"`
	BlockNumber          Optional[IntFilterInput]           `json:"blockNumber"`
	TxIndex              Optional[IntFilterInput]           `json:"txIndex"`
	Nonce                Optional[IntFilterInput]           `json:"nonce"`
	TxType               Optional[IntFilterInput]           `json:"txType"`
	LogsCount            Optional[IntFilterInput]           `json:"logsCount"`
	ValueWei             Optional[BigIntFilterInput]        `json:"valueWei"`
	Gas                  Optional[BigIntFilterInput]        `json:"gas"`
	GasPrice             Optional[BigIntFilterInput]        `json:"gasPrice"`
	GasUsed              Optional[BigIntFilterInput]        `json:"gasUsed"`
	MaxFeePerGas         Optional[BigIntFilterInput]        `json:"maxFeePerGas"`
	MaxPriorityFeePerGas Optional[BigIntFilterInput]        `json:"maxPriorityFeePerGas"`
	Success              Optional[IntFilterInput]           `json:"success"`
	Size                 Optional[IntFilterInput]           `json:"size"`
	ChainID              Optional[IntFilterInput]           `json:"chainId"`
	And                  Optional[[]TransactionFilterInput] `json:"and"`
	Or                   Optional[[]TransactionFilterInput] `json:"or"`
	Not                  Optional[TransactionFilterInput]   `json:"not"`
}

func (v TransactionFilterInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
