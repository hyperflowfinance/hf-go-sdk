package hyperflow

func (s *DelegateStakeChangedEventFields) Typename() *DelegateStakeChangedEventFields {
	s.node.field("__typename")
	return s
}

func (s *DelegateStakeChangedEventFields) Block(selectFields func(*BlockFields)) *DelegateStakeChangedEventFields {
	child := s.node.object("block")
	if selectFields != nil {
		selectFields(&BlockFields{node: child})
	}
	return s
}

func (s *DelegateStakeChangedEventFields) Transaction(selectFields func(*TransactionFields)) *DelegateStakeChangedEventFields {
	child := s.node.object("transaction")
	if selectFields != nil {
		selectFields(&TransactionFields{node: child})
	}
	return s
}

func (s *DelegateStakeChangedEventFields) Log(selectFields func(*LogFields)) *DelegateStakeChangedEventFields {
	child := s.node.object("log")
	if selectFields != nil {
		selectFields(&LogFields{node: child})
	}
	return s
}

func (s *DelegateStakeChangedEventFields) ContractAddress() *DelegateStakeChangedEventFields {
	s.node.field("contractAddress")
	return s
}

func (s *DelegateStakeChangedEventFields) Delegator() *DelegateStakeChangedEventFields {
	s.node.field("delegator")
	return s
}

func (s *DelegateStakeChangedEventFields) Guardian() *DelegateStakeChangedEventFields {
	s.node.field("guardian")
	return s
}

func (s *DelegateStakeChangedEventFields) SelfDelegatedStake(selectFields func(*StakeAmountFields)) *DelegateStakeChangedEventFields {
	child := s.node.object("selfDelegatedStake")
	if selectFields != nil {
		selectFields(&StakeAmountFields{node: child})
	}
	return s
}

func (s *DelegateStakeChangedEventFields) DelegatedStake(selectFields func(*StakeAmountFields)) *DelegateStakeChangedEventFields {
	child := s.node.object("delegatedStake")
	if selectFields != nil {
		selectFields(&StakeAmountFields{node: child})
	}
	return s
}

func (s *DelegateStakeChangedEventFields) DelegatorContributedStake(selectFields func(*StakeAmountFields)) *DelegateStakeChangedEventFields {
	child := s.node.object("delegatorContributedStake")
	if selectFields != nil {
		selectFields(&StakeAmountFields{node: child})
	}
	return s
}

func (s *DelegateStakeChangedEventFields) BlockNumber() *DelegateStakeChangedEventFields {
	s.node.field("blockNumber")
	return s
}

func (s *DelegateStakeChangedEventFields) BlockTimestamp() *DelegateStakeChangedEventFields {
	s.node.field("blockTimestamp")
	return s
}
