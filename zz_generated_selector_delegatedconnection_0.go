package hyperflow

func (s *DelegatedConnectionFields) Typename() *DelegatedConnectionFields {
	s.node.field("__typename")
	return s
}

func (s *DelegatedConnectionFields) Items(selectFields func(*DelegatedEventFields)) *DelegatedConnectionFields {
	child := s.node.object("items")
	if selectFields != nil {
		selectFields(&DelegatedEventFields{node: child})
	}
	return s
}

func (s *DelegatedConnectionFields) PageInfo(selectFields func(*PageInfoFields)) *DelegatedConnectionFields {
	child := s.node.object("pageInfo")
	if selectFields != nil {
		selectFields(&PageInfoFields{node: child})
	}
	return s
}
