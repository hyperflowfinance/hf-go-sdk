package hyperflow

func (s *Erc20TokenHolderFields) Typename() *Erc20TokenHolderFields {
	s.node.field("__typename")
	return s
}

func (s *Erc20TokenHolderFields) WalletAddress() *Erc20TokenHolderFields {
	s.node.field("walletAddress")
	return s
}

func (s *Erc20TokenHolderFields) ChainID() *Erc20TokenHolderFields {
	s.node.field("chainId")
	return s
}

func (s *Erc20TokenHolderFields) Balance(selectFields func(*BalanceResultFields)) *Erc20TokenHolderFields {
	child := s.node.object("balance")
	if selectFields != nil {
		selectFields(&BalanceResultFields{node: child})
	}
	return s
}
