package hyperflow

func (s *Erc20TokenTransferFields) BlockTimestamp() *Erc20TokenTransferFields {
	s.node.field("blockTimestamp")
	return s
}

func (s *Erc20TokenTransferFields) ParsedTransfer(selectFields func(*ParsedTokenAmountFields)) *Erc20TokenTransferFields {
	child := s.node.object("parsedTransfer")
	if selectFields != nil {
		selectFields(&ParsedTokenAmountFields{node: child})
	}
	return s
}
