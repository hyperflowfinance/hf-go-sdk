package hyperflow

type StorageAtInput struct {
	Address         string `json:"address"`
	Position        string `json:"position"`
	BlockIdentifier string `json:"blockIdentifier"`
}

func (v StorageAtInput) MarshalJSON() ([]byte, error) {
	return marshalInput(v)
}
