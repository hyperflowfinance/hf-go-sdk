package hyperflow

func (s *NodeMetadataFields) Typename() *NodeMetadataFields {
	s.node.field("__typename")
	return s
}

func (s *NodeMetadataFields) LastBlockNumber() *NodeMetadataFields {
	s.node.field("lastBlockNumber")
	return s
}

func (s *NodeMetadataFields) LastBlockHash() *NodeMetadataFields {
	s.node.field("lastBlockHash")
	return s
}

func (s *NodeMetadataFields) BlockTimestamp() *NodeMetadataFields {
	s.node.field("blockTimestamp")
	return s
}

func (s *NodeMetadataFields) NetworkID() *NodeMetadataFields {
	s.node.field("networkId")
	return s
}

func (s *NodeMetadataFields) ChainID() *NodeMetadataFields {
	s.node.field("chainId")
	return s
}

func (s *NodeMetadataFields) LastQueryTime() *NodeMetadataFields {
	s.node.field("lastQueryTime")
	return s
}
