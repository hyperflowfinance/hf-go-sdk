package hyperflow

func (s *Erc20TokenApprovalFields) Typename() *Erc20TokenApprovalFields {
	s.node.field("__typename")
	return s
}

func (s *Erc20TokenApprovalFields) Block(selectFields func(*BlockFields)) *Erc20TokenApprovalFields {
	child := s.node.object("block")
	if selectFields != nil {
		selectFields(&BlockFields{node: child})
	}
	return s
}

func (s *Erc20TokenApprovalFields) Transaction(selectFields func(*TransactionFields)) *Erc20TokenApprovalFields {
	child := s.node.object("transaction")
	if selectFields != nil {
		selectFields(&TransactionFields{node: child})
	}
	return s
}

func (s *Erc20TokenApprovalFields) Log(selectFields func(*LogFields)) *Erc20TokenApprovalFields {
	child := s.node.object("log")
	if selectFields != nil {
		selectFields(&LogFields{node: child})
	}
	return s
}

func (s *Erc20TokenApprovalFields) TokenAddress() *Erc20TokenApprovalFields {
	s.node.field("tokenAddress")
	return s
}

func (s *Erc20TokenApprovalFields) Owner() *Erc20TokenApprovalFields {
	s.node.field("owner")
	return s
}

func (s *Erc20TokenApprovalFields) Spender() *Erc20TokenApprovalFields {
	s.node.field("spender")
	return s
}

func (s *Erc20TokenApprovalFields) Value() *Erc20TokenApprovalFields {
	s.node.field("value")
	return s
}

func (s *Erc20TokenApprovalFields) BlockNumber() *Erc20TokenApprovalFields {
	s.node.field("blockNumber")
	return s
}

func (s *Erc20TokenApprovalFields) TxHash() *Erc20TokenApprovalFields {
	s.node.field("txHash")
	return s
}

func (s *Erc20TokenApprovalFields) LogIndex() *Erc20TokenApprovalFields {
	s.node.field("logIndex")
	return s
}

func (s *Erc20TokenApprovalFields) ChainID() *Erc20TokenApprovalFields {
	s.node.field("chainId")
	return s
}
