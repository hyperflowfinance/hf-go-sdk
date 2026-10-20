package hyperflow

func (s *TransactionCountResultFields) Typename() *TransactionCountResultFields {
	s.node.field("__typename")
	return s
}

func (s *TransactionCountResultFields) Hex() *TransactionCountResultFields {
	s.node.field("hex")
	return s
}

func (s *TransactionCountResultFields) Count() *TransactionCountResultFields {
	s.node.field("count")
	return s
}
