package hyperflow

func (s *LogConnectionFields) Typename() *LogConnectionFields {
	s.node.field("__typename")
	return s
}

func (s *LogConnectionFields) Items(selectFields func(*LogFields)) *LogConnectionFields {
	child := s.node.object("items")
	if selectFields != nil {
		selectFields(&LogFields{node: child})
	}
	return s
}

func (s *LogConnectionFields) PageInfo(selectFields func(*PageInfoFields)) *LogConnectionFields {
	child := s.node.object("pageInfo")
	if selectFields != nil {
		selectFields(&PageInfoFields{node: child})
	}
	return s
}
