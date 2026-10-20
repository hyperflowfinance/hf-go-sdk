package hyperflow

type BlockFilterInput struct {
	Hash             Optional[StringFilterInput]  `json:"hash"`
	ParentHash       Optional[StringFilterInput]  `json:"parentHash"`
	Miner            Optional[StringFilterInput]  `json:"miner"`
	Nonce            Optional[StringFilterInput]  `json:"nonce"`
	StateRoot        Optional[StringFilterInput]  `json:"stateRoot"`
	TxRoot           Optional[StringFilterInput]  `json:"txRoot"`
	ReceiptRoot      Optional[StringFilterInput]  `json:"receiptRoot"`
	MixDigest        Optional[StringFilterInput]  `json:"mixDigest"`
	UncleHash        Optional[StringFilterInput]  `json:"uncleHash"`
	ExtraData        Optional[StringFilterInput]  `json:"extraData"`
	LogsBloom        Optional[StringFilterInput]  `json:"logsBloom"`
	WithdrawalsHash  Optional[StringFilterInput]  `json:"withdrawalsHash"`
	ParentBeaconRoot Optional[StringFilterInput]  `json:"parentBeaconRoot"`
	RequestsHash     Optional[StringFilterInput]  `json:"requestsHash"`
	Number           Optional[IntFilterInput]     `json:"number"`
	GasLimit         Optional[BigIntFilterInput]  `json:"gasLimit"`
	GasUsed          Optional[BigIntFilterInput]  `json:"gasUsed"`
	Timestamp        Optional[IntFilterInput]     `json:"timestamp"`
	Difficulty       Optional[BigIntFilterInput]  `json:"difficulty"`
	TotalDifficulty  Optional[BigIntFilterInput]  `json:"totalDifficulty"`
	Size             Optional[IntFilterInput]     `json:"size"`
	BaseFee          Optional[BigIntFilterInput]  `json:"baseFee"`
	TxCount          Optional[IntFilterInput]     `json:"txCount"`
	UncleCount       Optional[IntFilterInput]     `json:"uncleCount"`
	WithdrawalCount  Optional[IntFilterInput]     `json:"withdrawalCount"`
	BlobGasUsed      Optional[BigIntFilterInput]  `json:"blobGasUsed"`
	ExcessBlobGas    Optional[BigIntFilterInput]  `json:"excessBlobGas"`
	ChainID          Optional[IntFilterInput]     `json:"chainId"`
	Finalized        Optional[BoolFilterInput]    `json:"finalized"`
	And              Optional[[]BlockFilterInput] `json:"and"`
	Or               Optional[[]BlockFilterInput] `json:"or"`
	Not              Optional[BlockFilterInput]   `json:"not"`
}

func (v BlockFilterInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
