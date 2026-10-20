package hyperflow

func (s *InternalTransactionFields) GasUsed() *InternalTransactionFields {
	s.node.field("gasUsed")
	return s
}

func (s *InternalTransactionFields) Input() *InternalTransactionFields {
	s.node.field("input")
	return s
}

func (s *InternalTransactionFields) Output() *InternalTransactionFields {
	s.node.field("output")
	return s
}

func (s *InternalTransactionFields) Error() *InternalTransactionFields {
	s.node.field("error")
	return s
}

func (s *InternalTransactionFields) Success() *InternalTransactionFields {
	s.node.field("success")
	return s
}

func (s *InternalTransactionFields) ChainID() *InternalTransactionFields {
	s.node.field("chainId")
	return s
}

func (s *InternalTransactionFields) Block(selectFields func(*BlockFields)) *InternalTransactionFields {
	child := s.node.object("block")
	if selectFields != nil {
		selectFields(&BlockFields{node: child})
	}
	return s
}

func (s *InternalTransactionFields) Transaction(selectFields func(*TransactionFields)) *InternalTransactionFields {
	child := s.node.object("transaction")
	if selectFields != nil {
		selectFields(&TransactionFields{node: child})
	}
	return s
}
