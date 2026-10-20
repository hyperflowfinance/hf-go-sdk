package hyperflow

func (s *TransactionResultFields) Typename() *TransactionResultFields {
	s.node.field("__typename")
	return s
}

func (s *TransactionResultFields) Hash() *TransactionResultFields {
	s.node.field("hash")
	return s
}

func (s *TransactionResultFields) Success() *TransactionResultFields {
	s.node.field("success")
	return s
}

func (s *TransactionResultFields) Error() *TransactionResultFields {
	s.node.field("error")
	return s
}
