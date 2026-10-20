package hyperflow

func (s *DelegatedEventFields) ChainID() *DelegatedEventFields {
	s.node.field("chainId")
	return s
}

func (s *DelegatedEventFields) Token(selectFields func(*Erc20TokenFields)) *DelegatedEventFields {
	child := s.node.object("token")
	if selectFields != nil {
		selectFields(&Erc20TokenFields{node: child})
	}
	return s
}
