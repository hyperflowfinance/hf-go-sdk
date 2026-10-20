package hyperflow

func (s *InternalTransactionConnectionFields) Typename() *InternalTransactionConnectionFields {
	s.node.field("__typename")
	return s
}

func (s *InternalTransactionConnectionFields) Items(selectFields func(*InternalTransactionFields)) *InternalTransactionConnectionFields {
	child := s.node.object("items")
	if selectFields != nil {
		selectFields(&InternalTransactionFields{node: child})
	}
	return s
}

func (s *InternalTransactionConnectionFields) PageInfo(selectFields func(*PageInfoFields)) *InternalTransactionConnectionFields {
	child := s.node.object("pageInfo")
	if selectFields != nil {
		selectFields(&PageInfoFields{node: child})
	}
	return s
}
