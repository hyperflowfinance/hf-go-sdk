package hyperflow

func (s *BlockFields) ExcessBlobGas() *BlockFields {
	s.node.field("excessBlobGas")
	return s
}

func (s *BlockFields) LogsBloom() *BlockFields {
	s.node.field("logsBloom")
	return s
}

func (s *BlockFields) ExtraData() *BlockFields {
	s.node.field("extraData")
	return s
}

func (s *BlockFields) ParentBeaconRoot() *BlockFields {
	s.node.field("parentBeaconRoot")
	return s
}

func (s *BlockFields) RequestsHash() *BlockFields {
	s.node.field("requestsHash")
	return s
}

func (s *BlockFields) ChainID() *BlockFields {
	s.node.field("chainId")
	return s
}

func (s *BlockFields) Finalized() *BlockFields {
	s.node.field("finalized")
	return s
}

func (s *BlockFields) Transactions(selectFields func(*TransactionFields)) *BlockFields {
	child := s.node.object("transactions")
	if selectFields != nil {
		selectFields(&TransactionFields{node: child})
	}
	return s
}
