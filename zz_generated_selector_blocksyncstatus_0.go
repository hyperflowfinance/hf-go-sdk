package hyperflow

func (s *BlockSyncStatusFields) Typename() *BlockSyncStatusFields {
	s.node.field("__typename")
	return s
}

func (s *BlockSyncStatusFields) NodeLastBlock() *BlockSyncStatusFields {
	s.node.field("nodeLastBlock")
	return s
}

func (s *BlockSyncStatusFields) DbMaxBlock() *BlockSyncStatusFields {
	s.node.field("dbMaxBlock")
	return s
}

func (s *BlockSyncStatusFields) DbActualCount() *BlockSyncStatusFields {
	s.node.field("dbActualCount")
	return s
}

func (s *BlockSyncStatusFields) Expected() *BlockSyncStatusFields {
	s.node.field("expected")
	return s
}

func (s *BlockSyncStatusFields) Actual() *BlockSyncStatusFields {
	s.node.field("actual")
	return s
}

func (s *BlockSyncStatusFields) IsSynced() *BlockSyncStatusFields {
	s.node.field("isSynced")
	return s
}

func (s *BlockSyncStatusFields) Status() *BlockSyncStatusFields {
	s.node.field("status")
	return s
}

func (s *BlockSyncStatusFields) Message() *BlockSyncStatusFields {
	s.node.field("message")
	return s
}
