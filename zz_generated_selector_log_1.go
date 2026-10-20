package hyperflow

func (s *LogFields) Removed() *LogFields {
	s.node.field("removed")
	return s
}

func (s *LogFields) ChainID() *LogFields {
	s.node.field("chainId")
	return s
}

func (s *LogFields) Topics() *LogFields {
	s.node.field("topics")
	return s
}

func (s *LogFields) Block(selectFields func(*BlockFields)) *LogFields {
	child := s.node.object("block")
	if selectFields != nil {
		selectFields(&BlockFields{node: child})
	}
	return s
}

func (s *LogFields) Transaction(selectFields func(*EmbeddedTransactionFields)) *LogFields {
	child := s.node.object("transaction")
	if selectFields != nil {
		selectFields(&EmbeddedTransactionFields{node: child})
	}
	return s
}
