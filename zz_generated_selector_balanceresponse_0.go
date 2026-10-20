package hyperflow

func (s *BalanceResponseFields) Typename() *BalanceResponseFields {
	s.node.field("__typename")
	return s
}

func (s *BalanceResponseFields) ChainID() *BalanceResponseFields {
	s.node.field("chainId")
	return s
}

func (s *BalanceResponseFields) Balance(selectFields func(*BalanceResultFields)) *BalanceResponseFields {
	child := s.node.object("balance")
	if selectFields != nil {
		selectFields(&BalanceResultFields{node: child})
	}
	return s
}
