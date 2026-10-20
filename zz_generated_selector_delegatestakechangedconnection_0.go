package hyperflow

func (s *DelegateStakeChangedConnectionFields) Typename() *DelegateStakeChangedConnectionFields {
	s.node.field("__typename")
	return s
}

func (s *DelegateStakeChangedConnectionFields) Items(selectFields func(*DelegateStakeChangedEventFields)) *DelegateStakeChangedConnectionFields {
	child := s.node.object("items")
	if selectFields != nil {
		selectFields(&DelegateStakeChangedEventFields{node: child})
	}
	return s
}

func (s *DelegateStakeChangedConnectionFields) PageInfo(selectFields func(*PageInfoFields)) *DelegateStakeChangedConnectionFields {
	child := s.node.object("pageInfo")
	if selectFields != nil {
		selectFields(&PageInfoFields{node: child})
	}
	return s
}
