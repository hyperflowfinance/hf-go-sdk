package hyperflow

func (s *Erc20TokenSwapEventConnectionFields) Typename() *Erc20TokenSwapEventConnectionFields {
	s.node.field("__typename")
	return s
}

func (s *Erc20TokenSwapEventConnectionFields) Items(selectFields func(*Erc20TokenSwapEventFields)) *Erc20TokenSwapEventConnectionFields {
	child := s.node.object("items")
	if selectFields != nil {
		selectFields(&Erc20TokenSwapEventFields{node: child})
	}
	return s
}

func (s *Erc20TokenSwapEventConnectionFields) PageInfo(selectFields func(*PageInfoFields)) *Erc20TokenSwapEventConnectionFields {
	child := s.node.object("pageInfo")
	if selectFields != nil {
		selectFields(&PageInfoFields{node: child})
	}
	return s
}
