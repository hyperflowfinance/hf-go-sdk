package hyperflow

func (s *ParsedSwapFields) Typename() *ParsedSwapFields {
	s.node.field("__typename")
	return s
}

func (s *ParsedSwapFields) TokenIn(selectFields func(*SwapTokenInfoFields)) *ParsedSwapFields {
	child := s.node.object("tokenIn")
	if selectFields != nil {
		selectFields(&SwapTokenInfoFields{node: child})
	}
	return s
}

func (s *ParsedSwapFields) TokenOut(selectFields func(*SwapTokenInfoFields)) *ParsedSwapFields {
	child := s.node.object("tokenOut")
	if selectFields != nil {
		selectFields(&SwapTokenInfoFields{node: child})
	}
	return s
}
