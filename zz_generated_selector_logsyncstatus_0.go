package hyperflow

func (s *LogSyncStatusFields) Typename() *LogSyncStatusFields {
	s.node.field("__typename")
	return s
}

func (s *LogSyncStatusFields) SumOfLogCountersInTxs() *LogSyncStatusFields {
	s.node.field("sumOfLogCountersInTxs")
	return s
}

func (s *LogSyncStatusFields) ActualCountInDb() *LogSyncStatusFields {
	s.node.field("actualCountInDb")
	return s
}

func (s *LogSyncStatusFields) IsSynced() *LogSyncStatusFields {
	s.node.field("isSynced")
	return s
}

func (s *LogSyncStatusFields) Status() *LogSyncStatusFields {
	s.node.field("status")
	return s
}

func (s *LogSyncStatusFields) Message() *LogSyncStatusFields {
	s.node.field("message")
	return s
}
