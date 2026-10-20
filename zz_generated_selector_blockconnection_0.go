package hyperflow

func (s *BlockConnectionFields) Typename() *BlockConnectionFields {
	s.node.field("__typename")
	return s
}

func (s *BlockConnectionFields) Items(selectFields func(*BlockFields)) *BlockConnectionFields {
	child := s.node.object("items")
	if selectFields != nil {
		selectFields(&BlockFields{node: child})
	}
	return s
}

func (s *BlockConnectionFields) PageInfo(selectFields func(*PageInfoFields)) *BlockConnectionFields {
	child := s.node.object("pageInfo")
	if selectFields != nil {
		selectFields(&PageInfoFields{node: child})
	}
	return s
}
