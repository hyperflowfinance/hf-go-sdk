package hyperflow

func (s *EmbeddedLogFields) Typename() *EmbeddedLogFields {
	s.node.field("__typename")
	return s
}

func (s *EmbeddedLogFields) BlockNumber() *EmbeddedLogFields {
	s.node.field("blockNumber")
	return s
}

func (s *EmbeddedLogFields) TxIndex() *EmbeddedLogFields {
	s.node.field("txIndex")
	return s
}

func (s *EmbeddedLogFields) LogIndex() *EmbeddedLogFields {
	s.node.field("logIndex")
	return s
}

func (s *EmbeddedLogFields) TxHash() *EmbeddedLogFields {
	s.node.field("txHash")
	return s
}

func (s *EmbeddedLogFields) BlockHash() *EmbeddedLogFields {
	s.node.field("blockHash")
	return s
}

func (s *EmbeddedLogFields) Address() *EmbeddedLogFields {
	s.node.field("address")
	return s
}

func (s *EmbeddedLogFields) Topic0() *EmbeddedLogFields {
	s.node.field("topic0")
	return s
}

func (s *EmbeddedLogFields) Topic1() *EmbeddedLogFields {
	s.node.field("topic1")
	return s
}

func (s *EmbeddedLogFields) Topic2() *EmbeddedLogFields {
	s.node.field("topic2")
	return s
}

func (s *EmbeddedLogFields) Topic3() *EmbeddedLogFields {
	s.node.field("topic3")
	return s
}

func (s *EmbeddedLogFields) Data() *EmbeddedLogFields {
	s.node.field("data")
	return s
}
