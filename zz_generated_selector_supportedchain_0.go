package hyperflow

func (s *SupportedChainFields) Typename() *SupportedChainFields {
	s.node.field("__typename")
	return s
}

func (s *SupportedChainFields) ChainID() *SupportedChainFields {
	s.node.field("chainId")
	return s
}

func (s *SupportedChainFields) ChainIDHex() *SupportedChainFields {
	s.node.field("chainIdHex")
	return s
}

func (s *SupportedChainFields) ChainName() *SupportedChainFields {
	s.node.field("chainName")
	return s
}

func (s *SupportedChainFields) NativeCurrencySymbol() *SupportedChainFields {
	s.node.field("nativeCurrencySymbol")
	return s
}

func (s *SupportedChainFields) NativeCurrencyDecimals() *SupportedChainFields {
	s.node.field("nativeCurrencyDecimals")
	return s
}

func (s *SupportedChainFields) DataMode() *SupportedChainFields {
	s.node.field("dataMode")
	return s
}
