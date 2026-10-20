package hyperflow

func (s *BlockRevertedEventFields) Typename() *BlockRevertedEventFields {
	s.node.field("__typename")
	return s
}

func (s *BlockRevertedEventFields) BlockNumber() *BlockRevertedEventFields {
	s.node.field("blockNumber")
	return s
}

func (s *BlockRevertedEventFields) BlockHash() *BlockRevertedEventFields {
	s.node.field("blockHash")
	return s
}
