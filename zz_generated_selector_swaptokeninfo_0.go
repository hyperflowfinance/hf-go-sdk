package hyperflow

func (s *SwapTokenInfoFields) Typename() *SwapTokenInfoFields {
	s.node.field("__typename")
	return s
}

func (s *SwapTokenInfoFields) Side() *SwapTokenInfoFields {
	s.node.field("side")
	return s
}

func (s *SwapTokenInfoFields) Address() *SwapTokenInfoFields {
	s.node.field("address")
	return s
}

func (s *SwapTokenInfoFields) Amount() *SwapTokenInfoFields {
	s.node.field("amount")
	return s
}

func (s *SwapTokenInfoFields) AmountRaw() *SwapTokenInfoFields {
	s.node.field("amountRaw")
	return s
}

func (s *SwapTokenInfoFields) Decimals() *SwapTokenInfoFields {
	s.node.field("decimals")
	return s
}

func (s *SwapTokenInfoFields) AmountFormatted() *SwapTokenInfoFields {
	s.node.field("amountFormatted")
	return s
}
