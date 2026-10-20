package hyperflow

type Erc20TokenSwapEventFilterInput struct {
	PairAddress    Optional[StringFilterInput]                `json:"pairAddress"`
	Sender         Optional[StringFilterInput]                `json:"sender"`
	To             Optional[StringFilterInput]                `json:"to"`
	TxHash         Optional[StringFilterInput]                `json:"txHash"`
	BlockNumber    Optional[IntFilterInput]                   `json:"blockNumber"`
	BlockTimestamp Optional[IntFilterInput]                   `json:"blockTimestamp"`
	LogIndex       Optional[IntFilterInput]                   `json:"logIndex"`
	ChainID        Optional[IntFilterInput]                   `json:"chainId"`
	And            Optional[[]Erc20TokenSwapEventFilterInput] `json:"and"`
	Or             Optional[[]Erc20TokenSwapEventFilterInput] `json:"or"`
	Not            Optional[Erc20TokenSwapEventFilterInput]   `json:"not"`
}

func (v Erc20TokenSwapEventFilterInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
