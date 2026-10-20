package hyperflow

func (s *Erc20TokenSwapEventFields) Typename() *Erc20TokenSwapEventFields {
	s.node.field("__typename")
	return s
}

func (s *Erc20TokenSwapEventFields) Block(selectFields func(*BlockFields)) *Erc20TokenSwapEventFields {
	child := s.node.object("block")
	if selectFields != nil {
		selectFields(&BlockFields{node: child})
	}
	return s
}

func (s *Erc20TokenSwapEventFields) Transaction(selectFields func(*TransactionFields)) *Erc20TokenSwapEventFields {
	child := s.node.object("transaction")
	if selectFields != nil {
		selectFields(&TransactionFields{node: child})
	}
	return s
}

func (s *Erc20TokenSwapEventFields) Log(selectFields func(*LogFields)) *Erc20TokenSwapEventFields {
	child := s.node.object("log")
	if selectFields != nil {
		selectFields(&LogFields{node: child})
	}
	return s
}

func (s *Erc20TokenSwapEventFields) PairAddress() *Erc20TokenSwapEventFields {
	s.node.field("pairAddress")
	return s
}

func (s *Erc20TokenSwapEventFields) Sender() *Erc20TokenSwapEventFields {
	s.node.field("sender")
	return s
}

func (s *Erc20TokenSwapEventFields) To() *Erc20TokenSwapEventFields {
	s.node.field("to")
	return s
}

func (s *Erc20TokenSwapEventFields) Amount0In() *Erc20TokenSwapEventFields {
	s.node.field("amount0In")
	return s
}

func (s *Erc20TokenSwapEventFields) Amount1In() *Erc20TokenSwapEventFields {
	s.node.field("amount1In")
	return s
}

func (s *Erc20TokenSwapEventFields) Amount0Out() *Erc20TokenSwapEventFields {
	s.node.field("amount0Out")
	return s
}

func (s *Erc20TokenSwapEventFields) Amount1Out() *Erc20TokenSwapEventFields {
	s.node.field("amount1Out")
	return s
}

func (s *Erc20TokenSwapEventFields) BlockNumber() *Erc20TokenSwapEventFields {
	s.node.field("blockNumber")
	return s
}
