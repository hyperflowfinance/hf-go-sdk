package hyperflow

func (s *TransactionFields) MaxFeePerGas() *TransactionFields {
	s.node.field("maxFeePerGas")
	return s
}

func (s *TransactionFields) MaxPriorityFeePerGas() *TransactionFields {
	s.node.field("maxPriorityFeePerGas")
	return s
}

func (s *TransactionFields) Input() *TransactionFields {
	s.node.field("input")
	return s
}

func (s *TransactionFields) V() *TransactionFields {
	s.node.field("v")
	return s
}

func (s *TransactionFields) R() *TransactionFields {
	s.node.field("r")
	return s
}

func (s *TransactionFields) S() *TransactionFields {
	s.node.field("s")
	return s
}

func (s *TransactionFields) Size() *TransactionFields {
	s.node.field("size")
	return s
}

func (s *TransactionFields) Success() *TransactionFields {
	s.node.field("success")
	return s
}

func (s *TransactionFields) LogsCount() *TransactionFields {
	s.node.field("logsCount")
	return s
}

func (s *TransactionFields) ChainID() *TransactionFields {
	s.node.field("chainId")
	return s
}

func (s *TransactionFields) Block(selectFields func(*BlockFields)) *TransactionFields {
	child := s.node.object("block")
	if selectFields != nil {
		selectFields(&BlockFields{node: child})
	}
	return s
}

func (s *TransactionFields) Logs(selectFields func(*EmbeddedLogFields)) *TransactionFields {
	child := s.node.object("logs")
	if selectFields != nil {
		selectFields(&EmbeddedLogFields{node: child})
	}
	return s
}
