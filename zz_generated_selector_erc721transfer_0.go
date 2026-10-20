package hyperflow

func (s *Erc721TransferFields) Typename() *Erc721TransferFields {
	s.node.field("__typename")
	return s
}

func (s *Erc721TransferFields) Block(selectFields func(*BlockFields)) *Erc721TransferFields {
	child := s.node.object("block")
	if selectFields != nil {
		selectFields(&BlockFields{node: child})
	}
	return s
}

func (s *Erc721TransferFields) Transaction(selectFields func(*TransactionFields)) *Erc721TransferFields {
	child := s.node.object("transaction")
	if selectFields != nil {
		selectFields(&TransactionFields{node: child})
	}
	return s
}

func (s *Erc721TransferFields) Log(selectFields func(*LogFields)) *Erc721TransferFields {
	child := s.node.object("log")
	if selectFields != nil {
		selectFields(&LogFields{node: child})
	}
	return s
}

func (s *Erc721TransferFields) TokenAddress() *Erc721TransferFields {
	s.node.field("tokenAddress")
	return s
}

func (s *Erc721TransferFields) FromAddress() *Erc721TransferFields {
	s.node.field("fromAddress")
	return s
}

func (s *Erc721TransferFields) ToAddress() *Erc721TransferFields {
	s.node.field("toAddress")
	return s
}

func (s *Erc721TransferFields) TokenID() *Erc721TransferFields {
	s.node.field("tokenId")
	return s
}

func (s *Erc721TransferFields) BlockNumber() *Erc721TransferFields {
	s.node.field("blockNumber")
	return s
}

func (s *Erc721TransferFields) TxHash() *Erc721TransferFields {
	s.node.field("txHash")
	return s
}

func (s *Erc721TransferFields) LogIndex() *Erc721TransferFields {
	s.node.field("logIndex")
	return s
}

func (s *Erc721TransferFields) ChainID() *Erc721TransferFields {
	s.node.field("chainId")
	return s
}
