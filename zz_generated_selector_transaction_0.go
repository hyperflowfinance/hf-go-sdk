package hyperflow

func (s *TransactionFields) Typename() *TransactionFields {
	s.node.field("__typename")
	return s
}

func (s *TransactionFields) BlockNumber() *TransactionFields {
	s.node.field("blockNumber")
	return s
}

func (s *TransactionFields) TxIndex() *TransactionFields {
	s.node.field("txIndex")
	return s
}

func (s *TransactionFields) Hash() *TransactionFields {
	s.node.field("hash")
	return s
}

func (s *TransactionFields) FromAddress() *TransactionFields {
	s.node.field("fromAddress")
	return s
}

func (s *TransactionFields) ToAddress() *TransactionFields {
	s.node.field("toAddress")
	return s
}

func (s *TransactionFields) ValueWei() *TransactionFields {
	s.node.field("valueWei")
	return s
}

func (s *TransactionFields) Gas() *TransactionFields {
	s.node.field("gas")
	return s
}

func (s *TransactionFields) GasPrice() *TransactionFields {
	s.node.field("gasPrice")
	return s
}

func (s *TransactionFields) GasUsed() *TransactionFields {
	s.node.field("gasUsed")
	return s
}

func (s *TransactionFields) Nonce() *TransactionFields {
	s.node.field("nonce")
	return s
}

func (s *TransactionFields) TxType() *TransactionFields {
	s.node.field("txType")
	return s
}
