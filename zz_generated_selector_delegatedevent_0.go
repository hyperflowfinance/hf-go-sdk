package hyperflow

func (s *DelegatedEventFields) Typename() *DelegatedEventFields {
	s.node.field("__typename")
	return s
}

func (s *DelegatedEventFields) Block(selectFields func(*BlockFields)) *DelegatedEventFields {
	child := s.node.object("block")
	if selectFields != nil {
		selectFields(&BlockFields{node: child})
	}
	return s
}

func (s *DelegatedEventFields) Transaction(selectFields func(*TransactionFields)) *DelegatedEventFields {
	child := s.node.object("transaction")
	if selectFields != nil {
		selectFields(&TransactionFields{node: child})
	}
	return s
}

func (s *DelegatedEventFields) Log(selectFields func(*LogFields)) *DelegatedEventFields {
	child := s.node.object("log")
	if selectFields != nil {
		selectFields(&LogFields{node: child})
	}
	return s
}

func (s *DelegatedEventFields) ContractAddress() *DelegatedEventFields {
	s.node.field("contractAddress")
	return s
}

func (s *DelegatedEventFields) Delegator() *DelegatedEventFields {
	s.node.field("delegator")
	return s
}

func (s *DelegatedEventFields) Guardian() *DelegatedEventFields {
	s.node.field("guardian")
	return s
}

func (s *DelegatedEventFields) BlockNumber() *DelegatedEventFields {
	s.node.field("blockNumber")
	return s
}

func (s *DelegatedEventFields) BlockTimestamp() *DelegatedEventFields {
	s.node.field("blockTimestamp")
	return s
}

func (s *DelegatedEventFields) TxHash() *DelegatedEventFields {
	s.node.field("txHash")
	return s
}

func (s *DelegatedEventFields) LogIndex() *DelegatedEventFields {
	s.node.field("logIndex")
	return s
}

func (s *DelegatedEventFields) TokenAddress() *DelegatedEventFields {
	s.node.field("tokenAddress")
	return s
}
