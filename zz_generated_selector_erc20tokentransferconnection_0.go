package hyperflow

func (s *Erc20TokenTransferConnectionFields) Typename() *Erc20TokenTransferConnectionFields {
	s.node.field("__typename")
	return s
}

func (s *Erc20TokenTransferConnectionFields) Items(selectFields func(*Erc20TokenTransferFields)) *Erc20TokenTransferConnectionFields {
	child := s.node.object("items")
	if selectFields != nil {
		selectFields(&Erc20TokenTransferFields{node: child})
	}
	return s
}

func (s *Erc20TokenTransferConnectionFields) PageInfo(selectFields func(*PageInfoFields)) *Erc20TokenTransferConnectionFields {
	child := s.node.object("pageInfo")
	if selectFields != nil {
		selectFields(&PageInfoFields{node: child})
	}
	return s
}
