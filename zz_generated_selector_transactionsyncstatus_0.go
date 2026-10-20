package hyperflow

func (s *TransactionSyncStatusFields) Typename() *TransactionSyncStatusFields {
	s.node.field("__typename")
	return s
}

func (s *TransactionSyncStatusFields) SumOfTxCountersInBlocks() *TransactionSyncStatusFields {
	s.node.field("sumOfTxCountersInBlocks")
	return s
}

func (s *TransactionSyncStatusFields) ActualCountInDb() *TransactionSyncStatusFields {
	s.node.field("actualCountInDb")
	return s
}

func (s *TransactionSyncStatusFields) IsSynced() *TransactionSyncStatusFields {
	s.node.field("isSynced")
	return s
}

func (s *TransactionSyncStatusFields) Status() *TransactionSyncStatusFields {
	s.node.field("status")
	return s
}

func (s *TransactionSyncStatusFields) Message() *TransactionSyncStatusFields {
	s.node.field("message")
	return s
}
