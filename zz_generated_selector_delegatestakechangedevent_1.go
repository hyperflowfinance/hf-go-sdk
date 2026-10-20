package hyperflow

func (s *DelegateStakeChangedEventFields) TxHash() *DelegateStakeChangedEventFields {
	s.node.field("txHash")
	return s
}

func (s *DelegateStakeChangedEventFields) LogIndex() *DelegateStakeChangedEventFields {
	s.node.field("logIndex")
	return s
}

func (s *DelegateStakeChangedEventFields) TokenAddress() *DelegateStakeChangedEventFields {
	s.node.field("tokenAddress")
	return s
}

func (s *DelegateStakeChangedEventFields) ChainID() *DelegateStakeChangedEventFields {
	s.node.field("chainId")
	return s
}

func (s *DelegateStakeChangedEventFields) Token(selectFields func(*Erc20TokenFields)) *DelegateStakeChangedEventFields {
	child := s.node.object("token")
	if selectFields != nil {
		selectFields(&Erc20TokenFields{node: child})
	}
	return s
}

func (s *DelegateStakeChangedEventFields) DelegatedEvents(selectFields func(*DelegatedEventFields)) *DelegateStakeChangedEventFields {
	child := s.node.object("delegatedEvents")
	if selectFields != nil {
		selectFields(&DelegatedEventFields{node: child})
	}
	return s
}
