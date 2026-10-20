package hyperflow

func (s *Erc721TransferConnectionFields) Typename() *Erc721TransferConnectionFields {
	s.node.field("__typename")
	return s
}

func (s *Erc721TransferConnectionFields) Items(selectFields func(*Erc721TransferFields)) *Erc721TransferConnectionFields {
	child := s.node.object("items")
	if selectFields != nil {
		selectFields(&Erc721TransferFields{node: child})
	}
	return s
}

func (s *Erc721TransferConnectionFields) PageInfo(selectFields func(*PageInfoFields)) *Erc721TransferConnectionFields {
	child := s.node.object("pageInfo")
	if selectFields != nil {
		selectFields(&PageInfoFields{node: child})
	}
	return s
}
