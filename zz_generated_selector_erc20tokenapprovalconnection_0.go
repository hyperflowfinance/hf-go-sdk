package hyperflow

func (s *Erc20TokenApprovalConnectionFields) Typename() *Erc20TokenApprovalConnectionFields {
	s.node.field("__typename")
	return s
}

func (s *Erc20TokenApprovalConnectionFields) Items(selectFields func(*Erc20TokenApprovalFields)) *Erc20TokenApprovalConnectionFields {
	child := s.node.object("items")
	if selectFields != nil {
		selectFields(&Erc20TokenApprovalFields{node: child})
	}
	return s
}

func (s *Erc20TokenApprovalConnectionFields) PageInfo(selectFields func(*PageInfoFields)) *Erc20TokenApprovalConnectionFields {
	child := s.node.object("pageInfo")
	if selectFields != nil {
		selectFields(&PageInfoFields{node: child})
	}
	return s
}
