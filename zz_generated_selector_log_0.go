package hyperflow

func (s *LogFields) Typename() *LogFields {
	s.node.field("__typename")
	return s
}

func (s *LogFields) BlockNumber() *LogFields {
	s.node.field("blockNumber")
	return s
}

func (s *LogFields) TxIndex() *LogFields {
	s.node.field("txIndex")
	return s
}

func (s *LogFields) LogIndex() *LogFields {
	s.node.field("logIndex")
	return s
}

func (s *LogFields) TxHash() *LogFields {
	s.node.field("txHash")
	return s
}

func (s *LogFields) BlockHash() *LogFields {
	s.node.field("blockHash")
	return s
}

func (s *LogFields) Address() *LogFields {
	s.node.field("address")
	return s
}

func (s *LogFields) Topic0() *LogFields {
	s.node.field("topic0")
	return s
}

func (s *LogFields) Topic1() *LogFields {
	s.node.field("topic1")
	return s
}

func (s *LogFields) Topic2() *LogFields {
	s.node.field("topic2")
	return s
}

func (s *LogFields) Topic3() *LogFields {
	s.node.field("topic3")
	return s
}

func (s *LogFields) Data() *LogFields {
	s.node.field("data")
	return s
}
