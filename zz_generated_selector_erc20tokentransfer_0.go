package hyperflow

func (s *Erc20TokenTransferFields) Typename() *Erc20TokenTransferFields {
	s.node.field("__typename")
	return s
}

func (s *Erc20TokenTransferFields) Block(selectFields func(*BlockFields)) *Erc20TokenTransferFields {
	child := s.node.object("block")
	if selectFields != nil {
		selectFields(&BlockFields{node: child})
	}
	return s
}

func (s *Erc20TokenTransferFields) Transaction(selectFields func(*TransactionFields)) *Erc20TokenTransferFields {
	child := s.node.object("transaction")
	if selectFields != nil {
		selectFields(&TransactionFields{node: child})
	}
	return s
}

func (s *Erc20TokenTransferFields) Log(selectFields func(*LogFields)) *Erc20TokenTransferFields {
	child := s.node.object("log")
	if selectFields != nil {
		selectFields(&LogFields{node: child})
	}
	return s
}

func (s *Erc20TokenTransferFields) TokenAddress() *Erc20TokenTransferFields {
	s.node.field("tokenAddress")
	return s
}

func (s *Erc20TokenTransferFields) FromAddress() *Erc20TokenTransferFields {
	s.node.field("fromAddress")
	return s
}

func (s *Erc20TokenTransferFields) ToAddress() *Erc20TokenTransferFields {
	s.node.field("toAddress")
	return s
}

func (s *Erc20TokenTransferFields) Value() *Erc20TokenTransferFields {
	s.node.field("value")
	return s
}

func (s *Erc20TokenTransferFields) BlockNumber() *Erc20TokenTransferFields {
	s.node.field("blockNumber")
	return s
}

func (s *Erc20TokenTransferFields) TxHash() *Erc20TokenTransferFields {
	s.node.field("txHash")
	return s
}

func (s *Erc20TokenTransferFields) LogIndex() *Erc20TokenTransferFields {
	s.node.field("logIndex")
	return s
}

func (s *Erc20TokenTransferFields) ChainID() *Erc20TokenTransferFields {
	s.node.field("chainId")
	return s
}
