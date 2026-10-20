package hyperflow

func (s *DbSyncStatusFields) Typename() *DbSyncStatusFields {
	s.node.field("__typename")
	return s
}

func (s *DbSyncStatusFields) Blocks(selectFields func(*BlockSyncStatusFields)) *DbSyncStatusFields {
	child := s.node.object("blocks")
	if selectFields != nil {
		selectFields(&BlockSyncStatusFields{node: child})
	}
	return s
}

func (s *DbSyncStatusFields) Transactions(selectFields func(*TransactionSyncStatusFields)) *DbSyncStatusFields {
	child := s.node.object("transactions")
	if selectFields != nil {
		selectFields(&TransactionSyncStatusFields{node: child})
	}
	return s
}

func (s *DbSyncStatusFields) Logs(selectFields func(*LogSyncStatusFields)) *DbSyncStatusFields {
	child := s.node.object("logs")
	if selectFields != nil {
		selectFields(&LogSyncStatusFields{node: child})
	}
	return s
}

func (s *DbSyncStatusFields) OverallStatus() *DbSyncStatusFields {
	s.node.field("overallStatus")
	return s
}

func (s *DbSyncStatusFields) SyncPercentage() *DbSyncStatusFields {
	s.node.field("syncPercentage")
	return s
}

func (s *DbSyncStatusFields) SuggestedAction() *DbSyncStatusFields {
	s.node.field("suggestedAction")
	return s
}

func (s *DbSyncStatusFields) LastSyncCheck() *DbSyncStatusFields {
	s.node.field("lastSyncCheck")
	return s
}

func (s *DbSyncStatusFields) DiagnosticMessage() *DbSyncStatusFields {
	s.node.field("diagnosticMessage")
	return s
}
