package hyperflow

func (s *Erc20TokenSwapEventFields) TxHash() *Erc20TokenSwapEventFields {
	s.node.field("txHash")
	return s
}

func (s *Erc20TokenSwapEventFields) LogIndex() *Erc20TokenSwapEventFields {
	s.node.field("logIndex")
	return s
}

func (s *Erc20TokenSwapEventFields) ChainID() *Erc20TokenSwapEventFields {
	s.node.field("chainId")
	return s
}

func (s *Erc20TokenSwapEventFields) BlockTimestamp() *Erc20TokenSwapEventFields {
	s.node.field("blockTimestamp")
	return s
}

func (s *Erc20TokenSwapEventFields) ParsedSwap(selectFields func(*ParsedSwapFields)) *Erc20TokenSwapEventFields {
	child := s.node.object("parsedSwap")
	if selectFields != nil {
		selectFields(&ParsedSwapFields{node: child})
	}
	return s
}
