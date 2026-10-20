package hyperflow

func (s *InternalTransactionFields) Typename() *InternalTransactionFields {
	s.node.field("__typename")
	return s
}

func (s *InternalTransactionFields) BlockNumber() *InternalTransactionFields {
	s.node.field("blockNumber")
	return s
}

func (s *InternalTransactionFields) TxHash() *InternalTransactionFields {
	s.node.field("txHash")
	return s
}

func (s *InternalTransactionFields) TxIndex() *InternalTransactionFields {
	s.node.field("txIndex")
	return s
}

func (s *InternalTransactionFields) TraceIndex() *InternalTransactionFields {
	s.node.field("traceIndex")
	return s
}

func (s *InternalTransactionFields) Subtraces() *InternalTransactionFields {
	s.node.field("subtraces")
	return s
}

func (s *InternalTransactionFields) TraceAddress() *InternalTransactionFields {
	s.node.field("traceAddress")
	return s
}

func (s *InternalTransactionFields) FromAddress() *InternalTransactionFields {
	s.node.field("fromAddress")
	return s
}

func (s *InternalTransactionFields) ToAddress() *InternalTransactionFields {
	s.node.field("toAddress")
	return s
}

func (s *InternalTransactionFields) ValueWei() *InternalTransactionFields {
	s.node.field("valueWei")
	return s
}

func (s *InternalTransactionFields) CallType() *InternalTransactionFields {
	s.node.field("callType")
	return s
}

func (s *InternalTransactionFields) Gas() *InternalTransactionFields {
	s.node.field("gas")
	return s
}
