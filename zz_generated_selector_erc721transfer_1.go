package hyperflow

func (s *Erc721TransferFields) BlockTimestamp() *Erc721TransferFields {
	s.node.field("blockTimestamp")
	return s
}
