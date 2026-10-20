package hyperflow

type Erc721TransferOrderByInput struct {
	Field     Erc721TransferOrderField `json:"field"`
	Direction SortDirection            `json:"direction"`
}

func (v Erc721TransferOrderByInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
