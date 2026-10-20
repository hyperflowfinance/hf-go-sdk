package hyperflow

func (s *Erc20TokenFields) Typename() *Erc20TokenFields {
	s.node.field("__typename")
	return s
}

func (s *Erc20TokenFields) TokenAddress() *Erc20TokenFields {
	s.node.field("tokenAddress")
	return s
}

func (s *Erc20TokenFields) Erc20TotalSupply(argChainID Optional[int32], selectFields func(*BalanceResponseFields)) *Erc20TokenFields {
	child := s.node.object("erc20TotalSupply", argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	if selectFields != nil {
		selectFields(&BalanceResponseFields{node: child})
	}
	return s
}

func (s *Erc20TokenFields) Erc20BalanceOf(argWalletAddress string, argChainID Optional[int32], selectFields func(*BalanceResponseFields)) *Erc20TokenFields {
	child := s.node.object("erc20BalanceOf", argument{name: "walletAddress", gqlType: "String!", value: argWalletAddress, optional: false}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	if selectFields != nil {
		selectFields(&BalanceResponseFields{node: child})
	}
	return s
}

func (s *Erc20TokenFields) Erc20Holders(argMinBalance string, argLimit int32, argOffset int32, argChainID Optional[int32], selectFields func(*Erc20TokenHolderFields)) *Erc20TokenFields {
	child := s.node.object("erc20Holders", argument{name: "minBalance", gqlType: "String!", value: argMinBalance, optional: false}, argument{name: "limit", gqlType: "Int!", value: argLimit, optional: false}, argument{name: "offset", gqlType: "Int!", value: argOffset, optional: false}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	if selectFields != nil {
		selectFields(&Erc20TokenHolderFields{node: child})
	}
	return s
}
