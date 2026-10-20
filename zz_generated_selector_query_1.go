package hyperflow

func (s *QueryFields) Erc20TokenSwapEvents(argInput Optional[Erc20TokenSwapEventQueryInput], selectFields func(*Erc20TokenSwapEventConnectionFields)) *QueryFields {
	child := s.node.object("erc20TokenSwapEvents", argument{name: "input", gqlType: "ERC20TokenSwapEventQueryInput", value: argInput, optional: true})
	if selectFields != nil {
		selectFields(&Erc20TokenSwapEventConnectionFields{node: child})
	}
	return s
}

func (s *QueryFields) Erc721Transfers(argInput Optional[Erc721TransferQueryInput], selectFields func(*Erc721TransferConnectionFields)) *QueryFields {
	child := s.node.object("erc721Transfers", argument{name: "input", gqlType: "ERC721TransferQueryInput", value: argInput, optional: true})
	if selectFields != nil {
		selectFields(&Erc721TransferConnectionFields{node: child})
	}
	return s
}

func (s *QueryFields) UsageStat(selectFields func(*EffectivenessFields)) *QueryFields {
	child := s.node.object("usageStat")
	if selectFields != nil {
		selectFields(&EffectivenessFields{node: child})
	}
	return s
}

func (s *QueryFields) GetStorageAt(argInput StorageAtInput, argChain Optional[string], argChainID Optional[int32], selectFields func(*StorageAtResultFields)) *QueryFields {
	child := s.node.object("getStorageAt", argument{name: "input", gqlType: "StorageAtInput!", value: argInput, optional: false}, argument{name: "chain", gqlType: "String", value: argChain, optional: true}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	if selectFields != nil {
		selectFields(&StorageAtResultFields{node: child})
	}
	return s
}

func (s *QueryFields) EstimateGas(argInput EstimateGasInput, argChain Optional[string], argChainID Optional[int32], selectFields func(*EstimateGasResultFields)) *QueryFields {
	child := s.node.object("estimateGas", argument{name: "input", gqlType: "EstimateGasInput!", value: argInput, optional: false}, argument{name: "chain", gqlType: "String", value: argChain, optional: true}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	if selectFields != nil {
		selectFields(&EstimateGasResultFields{node: child})
	}
	return s
}

func (s *QueryFields) Balance(argWalletAddress string, argBlockNumber Optional[int32], argChainID Optional[int32], selectFields func(*BalanceResponseFields)) *QueryFields {
	child := s.node.object("balance", argument{name: "walletAddress", gqlType: "String!", value: argWalletAddress, optional: false}, argument{name: "blockNumber", gqlType: "Int", value: argBlockNumber, optional: true}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	if selectFields != nil {
		selectFields(&BalanceResponseFields{node: child})
	}
	return s
}

func (s *QueryFields) EthCall(argInput EthCallInput, argChain Optional[string], argChainID Optional[int32], selectFields func(*EthCallResultFields)) *QueryFields {
	child := s.node.object("ethCall", argument{name: "input", gqlType: "EthCallInput!", value: argInput, optional: false}, argument{name: "chain", gqlType: "String", value: argChain, optional: true}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	if selectFields != nil {
		selectFields(&EthCallResultFields{node: child})
	}
	return s
}

func (s *QueryFields) EthGetTransactionCount(argInput GetTransactionCountInput, selectFields func(*TransactionCountResultFields)) *QueryFields {
	child := s.node.object("ethGetTransactionCount", argument{name: "input", gqlType: "GetTransactionCountInput!", value: argInput, optional: false})
	if selectFields != nil {
		selectFields(&TransactionCountResultFields{node: child})
	}
	return s
}

func (s *QueryFields) DbSyncStatus(argChain Optional[string], argChainID Optional[string], selectFields func(*DbSyncResponseFields)) *QueryFields {
	child := s.node.object("dbSyncStatus", argument{name: "chain", gqlType: "String", value: argChain, optional: true}, argument{name: "chainId", gqlType: "String", value: argChainID, optional: true})
	if selectFields != nil {
		selectFields(&DbSyncResponseFields{node: child})
	}
	return s
}

func (s *QueryFields) DelegateStakeChangedEvents(argInput Optional[DelegateStakeChangedQueryInput], selectFields func(*DelegateStakeChangedConnectionFields)) *QueryFields {
	child := s.node.object("delegateStakeChangedEvents", argument{name: "input", gqlType: "DelegateStakeChangedQueryInput", value: argInput, optional: true})
	if selectFields != nil {
		selectFields(&DelegateStakeChangedConnectionFields{node: child})
	}
	return s
}

func (s *QueryFields) DelegatedEvents(argInput Optional[DelegatedQueryInput], selectFields func(*DelegatedConnectionFields)) *QueryFields {
	child := s.node.object("delegatedEvents", argument{name: "input", gqlType: "DelegatedQueryInput", value: argInput, optional: true})
	if selectFields != nil {
		selectFields(&DelegatedConnectionFields{node: child})
	}
	return s
}

func (s *QueryFields) SupportedChains(selectFields func(*SupportedChainFields)) *QueryFields {
	child := s.node.object("supportedChains")
	if selectFields != nil {
		selectFields(&SupportedChainFields{node: child})
	}
	return s
}
