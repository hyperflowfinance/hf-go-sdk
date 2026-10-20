package hyperflow

type LogFilterInput struct {
	BlockNumber Optional[IntFilterInput]    `json:"blockNumber"`
	TxIndex     Optional[IntFilterInput]    `json:"txIndex"`
	LogIndex    Optional[IntFilterInput]    `json:"logIndex"`
	TxHash      Optional[StringFilterInput] `json:"txHash"`
	BlockHash   Optional[StringFilterInput] `json:"blockHash"`
	Address     Optional[StringFilterInput] `json:"address"`
	Topic0      Optional[StringFilterInput] `json:"topic0"`
	Topic1      Optional[StringFilterInput] `json:"topic1"`
	Topic2      Optional[StringFilterInput] `json:"topic2"`
	Topic3      Optional[StringFilterInput] `json:"topic3"`
	Data        Optional[StringFilterInput] `json:"data"`
	Removed     Optional[BoolFilterInput]   `json:"removed"`
	ChainID     Optional[IntFilterInput]    `json:"chainId"`
	And         Optional[[]LogFilterInput]  `json:"and"`
	Or          Optional[[]LogFilterInput]  `json:"or"`
	Not         Optional[LogFilterInput]    `json:"not"`
}

func (v LogFilterInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
