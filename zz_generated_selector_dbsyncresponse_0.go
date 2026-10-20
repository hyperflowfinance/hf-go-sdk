package hyperflow

func (s *DbSyncResponseFields) Typename() *DbSyncResponseFields {
	s.node.field("__typename")
	return s
}

func (s *DbSyncResponseFields) NodeMetadata(selectFields func(*NodeMetadataFields)) *DbSyncResponseFields {
	child := s.node.object("nodeMetadata")
	if selectFields != nil {
		selectFields(&NodeMetadataFields{node: child})
	}
	return s
}

func (s *DbSyncResponseFields) DbSyncStatus(selectFields func(*DbSyncStatusFields)) *DbSyncResponseFields {
	child := s.node.object("dbSyncStatus")
	if selectFields != nil {
		selectFields(&DbSyncStatusFields{node: child})
	}
	return s
}
