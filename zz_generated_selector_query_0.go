package hyperflow

func (s *QueryFields) Typename() *QueryFields {
	s.node.field("__typename")
	return s
}

func (s *QueryFields) Metadata(selectFields func(*MetadataFields)) *QueryFields {
	child := s.node.object("metadata")
	if selectFields != nil {
		selectFields(&MetadataFields{node: child})
	}
	return s
}

func (s *QueryFields) EthSyncing(argChain Optional[string], argChainID Optional[int32]) *QueryFields {
	s.node.field("ethSyncing", argument{name: "chain", gqlType: "String", value: argChain, optional: true}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	return s
}

func (s *QueryFields) EthGetCode(argAddress string, argBlockNumber Optional[int32], argChainID Optional[int32]) *QueryFields {
	s.node.field("ethGetCode", argument{name: "address", gqlType: "String!", value: argAddress, optional: false}, argument{name: "blockNumber", gqlType: "Int", value: argBlockNumber, optional: true}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	return s
}

func (s *QueryFields) Web3Sha3(argMessage string) *QueryFields {
	s.node.field("web3Sha3", argument{name: "message", gqlType: "String!", value: argMessage, optional: false})
	return s
}

func (s *QueryFields) Transactions(argInput Optional[TransactionQueryInput], selectFields func(*TransactionConnectionFields)) *QueryFields {
	child := s.node.object("transactions", argument{name: "input", gqlType: "TransactionQueryInput", value: argInput, optional: true})
	if selectFields != nil {
		selectFields(&TransactionConnectionFields{node: child})
	}
	return s
}

func (s *QueryFields) Blocks(argQueryInput Optional[BlockQueryInput], selectFields func(*BlockConnectionFields)) *QueryFields {
	child := s.node.object("blocks", argument{name: "queryInput", gqlType: "BlockQueryInput", value: argQueryInput, optional: true})
	if selectFields != nil {
		selectFields(&BlockConnectionFields{node: child})
	}
	return s
}

func (s *QueryFields) Logs(argInput Optional[LogQueryInput], selectFields func(*LogConnectionFields)) *QueryFields {
	child := s.node.object("logs", argument{name: "input", gqlType: "LogQueryInput", value: argInput, optional: true})
	if selectFields != nil {
		selectFields(&LogConnectionFields{node: child})
	}
	return s
}

func (s *QueryFields) InternalTransactions(argInput Optional[InternalTxQueryInput], selectFields func(*InternalTransactionConnectionFields)) *QueryFields {
	child := s.node.object("internalTransactions", argument{name: "input", gqlType: "InternalTxQueryInput", value: argInput, optional: true})
	if selectFields != nil {
		selectFields(&InternalTransactionConnectionFields{node: child})
	}
	return s
}

func (s *QueryFields) Erc20Token(argTokenAddress string, selectFields func(*Erc20TokenFields)) *QueryFields {
	child := s.node.object("erc20Token", argument{name: "tokenAddress", gqlType: "String!", value: argTokenAddress, optional: false})
	if selectFields != nil {
		selectFields(&Erc20TokenFields{node: child})
	}
	return s
}

func (s *QueryFields) Erc20TokenTransfers(argInput Optional[Erc20TokenTransferQueryInput], selectFields func(*Erc20TokenTransferConnectionFields)) *QueryFields {
	child := s.node.object("erc20TokenTransfers", argument{name: "input", gqlType: "ERC20TokenTransferQueryInput", value: argInput, optional: true})
	if selectFields != nil {
		selectFields(&Erc20TokenTransferConnectionFields{node: child})
	}
	return s
}

func (s *QueryFields) Erc20TokenApprovals(argInput Optional[Erc20TokenApprovalQueryInput], selectFields func(*Erc20TokenApprovalConnectionFields)) *QueryFields {
	child := s.node.object("erc20TokenApprovals", argument{name: "input", gqlType: "ERC20TokenApprovalQueryInput", value: argInput, optional: true})
	if selectFields != nil {
		selectFields(&Erc20TokenApprovalConnectionFields{node: child})
	}
	return s
}
