package hyperflow

func (s *TransactionConnectionFields) Typename() *TransactionConnectionFields {
	s.node.field("__typename")
	return s
}

func (s *TransactionConnectionFields) Items(selectFields func(*TransactionFields)) *TransactionConnectionFields {
	child := s.node.object("items")
	if selectFields != nil {
		selectFields(&TransactionFields{node: child})
	}
	return s
}

func (s *TransactionConnectionFields) PageInfo(selectFields func(*PageInfoFields)) *TransactionConnectionFields {
	child := s.node.object("pageInfo")
	if selectFields != nil {
		selectFields(&PageInfoFields{node: child})
	}
	return s
}
