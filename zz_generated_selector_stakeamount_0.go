package hyperflow

func (s *StakeAmountFields) Typename() *StakeAmountFields {
	s.node.field("__typename")
	return s
}

func (s *StakeAmountFields) Hex() *StakeAmountFields {
	s.node.field("hex")
	return s
}

func (s *StakeAmountFields) Wei() *StakeAmountFields {
	s.node.field("wei")
	return s
}

func (s *StakeAmountFields) Decimals() *StakeAmountFields {
	s.node.field("decimals")
	return s
}

func (s *StakeAmountFields) Decimal() *StakeAmountFields {
	s.node.field("decimal")
	return s
}
