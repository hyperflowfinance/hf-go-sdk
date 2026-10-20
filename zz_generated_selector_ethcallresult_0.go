package hyperflow

func (s *EthCallResultFields) Typename() *EthCallResultFields {
	s.node.field("__typename")
	return s
}

func (s *EthCallResultFields) Data() *EthCallResultFields {
	s.node.field("data")
	return s
}
