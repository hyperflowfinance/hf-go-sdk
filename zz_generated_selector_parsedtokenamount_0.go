package hyperflow

func (s *ParsedTokenAmountFields) Typename() *ParsedTokenAmountFields {
	s.node.field("__typename")
	return s
}

func (s *ParsedTokenAmountFields) TokenAddress() *ParsedTokenAmountFields {
	s.node.field("tokenAddress")
	return s
}

func (s *ParsedTokenAmountFields) Decimals() *ParsedTokenAmountFields {
	s.node.field("decimals")
	return s
}

func (s *ParsedTokenAmountFields) Amount() *ParsedTokenAmountFields {
	s.node.field("amount")
	return s
}

func (s *ParsedTokenAmountFields) AmountRaw() *ParsedTokenAmountFields {
	s.node.field("amountRaw")
	return s
}

func (s *ParsedTokenAmountFields) AmountFormatted() *ParsedTokenAmountFields {
	s.node.field("amountFormatted")
	return s
}
