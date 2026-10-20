package hyperflow

func (s *EmbeddedLogFields) Removed() *EmbeddedLogFields {
	s.node.field("removed")
	return s
}

func (s *EmbeddedLogFields) ChainID() *EmbeddedLogFields {
	s.node.field("chainId")
	return s
}

func (s *EmbeddedLogFields) Topics() *EmbeddedLogFields {
	s.node.field("topics")
	return s
}
