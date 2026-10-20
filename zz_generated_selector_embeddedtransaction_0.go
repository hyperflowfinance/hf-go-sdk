package hyperflow

func (s *EmbeddedTransactionFields) Typename() *EmbeddedTransactionFields {
	s.node.field("__typename")
	return s
}

func (s *EmbeddedTransactionFields) BlockNumber() *EmbeddedTransactionFields {
	s.node.field("blockNumber")
	return s
}

func (s *EmbeddedTransactionFields) TxIndex() *EmbeddedTransactionFields {
	s.node.field("txIndex")
	return s
}

func (s *EmbeddedTransactionFields) Hash() *EmbeddedTransactionFields {
	s.node.field("hash")
	return s
}

func (s *EmbeddedTransactionFields) FromAddress() *EmbeddedTransactionFields {
	s.node.field("fromAddress")
	return s
}

func (s *EmbeddedTransactionFields) ToAddress() *EmbeddedTransactionFields {
	s.node.field("toAddress")
	return s
}

func (s *EmbeddedTransactionFields) ValueWei() *EmbeddedTransactionFields {
	s.node.field("valueWei")
	return s
}

func (s *EmbeddedTransactionFields) Gas() *EmbeddedTransactionFields {
	s.node.field("gas")
	return s
}

func (s *EmbeddedTransactionFields) GasPrice() *EmbeddedTransactionFields {
	s.node.field("gasPrice")
	return s
}

func (s *EmbeddedTransactionFields) GasUsed() *EmbeddedTransactionFields {
	s.node.field("gasUsed")
	return s
}

func (s *EmbeddedTransactionFields) Nonce() *EmbeddedTransactionFields {
	s.node.field("nonce")
	return s
}

func (s *EmbeddedTransactionFields) TxType() *EmbeddedTransactionFields {
	s.node.field("txType")
	return s
}
