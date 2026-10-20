package hyperflow

func (s *EmbeddedTransactionFields) MaxFeePerGas() *EmbeddedTransactionFields {
	s.node.field("maxFeePerGas")
	return s
}

func (s *EmbeddedTransactionFields) MaxPriorityFeePerGas() *EmbeddedTransactionFields {
	s.node.field("maxPriorityFeePerGas")
	return s
}

func (s *EmbeddedTransactionFields) Input() *EmbeddedTransactionFields {
	s.node.field("input")
	return s
}

func (s *EmbeddedTransactionFields) V() *EmbeddedTransactionFields {
	s.node.field("v")
	return s
}

func (s *EmbeddedTransactionFields) R() *EmbeddedTransactionFields {
	s.node.field("r")
	return s
}

func (s *EmbeddedTransactionFields) S() *EmbeddedTransactionFields {
	s.node.field("s")
	return s
}

func (s *EmbeddedTransactionFields) Size() *EmbeddedTransactionFields {
	s.node.field("size")
	return s
}

func (s *EmbeddedTransactionFields) Success() *EmbeddedTransactionFields {
	s.node.field("success")
	return s
}

func (s *EmbeddedTransactionFields) LogsCount() *EmbeddedTransactionFields {
	s.node.field("logsCount")
	return s
}

func (s *EmbeddedTransactionFields) ChainID() *EmbeddedTransactionFields {
	s.node.field("chainId")
	return s
}
