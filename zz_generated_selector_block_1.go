package hyperflow

func (s *BlockFields) TxCount() *BlockFields {
	s.node.field("txCount")
	return s
}

func (s *BlockFields) StateRoot() *BlockFields {
	s.node.field("stateRoot")
	return s
}

func (s *BlockFields) TxRoot() *BlockFields {
	s.node.field("txRoot")
	return s
}

func (s *BlockFields) ReceiptRoot() *BlockFields {
	s.node.field("receiptRoot")
	return s
}

func (s *BlockFields) Nonce() *BlockFields {
	s.node.field("nonce")
	return s
}

func (s *BlockFields) MixDigest() *BlockFields {
	s.node.field("mixDigest")
	return s
}

func (s *BlockFields) UncleHash() *BlockFields {
	s.node.field("uncleHash")
	return s
}

func (s *BlockFields) Uncles() *BlockFields {
	s.node.field("uncles")
	return s
}

func (s *BlockFields) UncleCount() *BlockFields {
	s.node.field("uncleCount")
	return s
}

func (s *BlockFields) WithdrawalsHash() *BlockFields {
	s.node.field("withdrawalsHash")
	return s
}

func (s *BlockFields) WithdrawalCount() *BlockFields {
	s.node.field("withdrawalCount")
	return s
}

func (s *BlockFields) BlobGasUsed() *BlockFields {
	s.node.field("blobGasUsed")
	return s
}
