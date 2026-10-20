package hyperflow

func (s *BalanceResultFields) Typename() *BalanceResultFields {
	s.node.field("__typename")
	return s
}

func (s *BalanceResultFields) Hex() *BalanceResultFields {
	s.node.field("hex")
	return s
}

func (s *BalanceResultFields) Wei() *BalanceResultFields {
	s.node.field("wei")
	return s
}

func (s *BalanceResultFields) Decimals() *BalanceResultFields {
	s.node.field("decimals")
	return s
}

func (s *BalanceResultFields) Decimal() *BalanceResultFields {
	s.node.field("decimal")
	return s
}
