package hyperflow

func (s *BlockFields) Typename() *BlockFields {
	s.node.field("__typename")
	return s
}

func (s *BlockFields) Number() *BlockFields {
	s.node.field("number")
	return s
}

func (s *BlockFields) Hash() *BlockFields {
	s.node.field("hash")
	return s
}

func (s *BlockFields) ParentHash() *BlockFields {
	s.node.field("parentHash")
	return s
}

func (s *BlockFields) Timestamp() *BlockFields {
	s.node.field("timestamp")
	return s
}

func (s *BlockFields) Miner() *BlockFields {
	s.node.field("miner")
	return s
}

func (s *BlockFields) GasUsed() *BlockFields {
	s.node.field("gasUsed")
	return s
}

func (s *BlockFields) GasLimit() *BlockFields {
	s.node.field("gasLimit")
	return s
}

func (s *BlockFields) BaseFee() *BlockFields {
	s.node.field("baseFee")
	return s
}

func (s *BlockFields) Difficulty() *BlockFields {
	s.node.field("difficulty")
	return s
}

func (s *BlockFields) TotalDifficulty() *BlockFields {
	s.node.field("totalDifficulty")
	return s
}

func (s *BlockFields) Size() *BlockFields {
	s.node.field("size")
	return s
}
