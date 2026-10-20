package hyperflow

func (s *SubscriptionFields) Typename() *SubscriptionFields {
	s.node.field("__typename")
	return s
}

func (s *SubscriptionFields) NewBlocks(argCriteria Optional[BlockFilterInput], selectFields func(*BlockFields)) *SubscriptionFields {
	child := s.node.object("newBlocks", argument{name: "criteria", gqlType: "BlockFilterInput", value: argCriteria, optional: true})
	if selectFields != nil {
		selectFields(&BlockFields{node: child})
	}
	return s
}

func (s *SubscriptionFields) NewLogs(argCriteria Optional[LogFilterInput], selectFields func(*LogFields)) *SubscriptionFields {
	child := s.node.object("newLogs", argument{name: "criteria", gqlType: "LogFilterInput", value: argCriteria, optional: true})
	if selectFields != nil {
		selectFields(&LogFields{node: child})
	}
	return s
}

func (s *SubscriptionFields) NewTransactions(argCriteria Optional[TransactionFilterInput], selectFields func(*TransactionFields)) *SubscriptionFields {
	child := s.node.object("newTransactions", argument{name: "criteria", gqlType: "TransactionFilterInput", value: argCriteria, optional: true})
	if selectFields != nil {
		selectFields(&TransactionFields{node: child})
	}
	return s
}

func (s *SubscriptionFields) NewErc20Transfers(argCriteria Optional[Erc20TokenTransferFilterInput], selectFields func(*Erc20TokenTransferFields)) *SubscriptionFields {
	child := s.node.object("newErc20Transfers", argument{name: "criteria", gqlType: "ERC20TokenTransferFilterInput", value: argCriteria, optional: true})
	if selectFields != nil {
		selectFields(&Erc20TokenTransferFields{node: child})
	}
	return s
}

func (s *SubscriptionFields) NewErc20Approvals(argCriteria Optional[Erc20TokenApprovalFilterInput], selectFields func(*Erc20TokenApprovalFields)) *SubscriptionFields {
	child := s.node.object("newErc20Approvals", argument{name: "criteria", gqlType: "ERC20TokenApprovalFilterInput", value: argCriteria, optional: true})
	if selectFields != nil {
		selectFields(&Erc20TokenApprovalFields{node: child})
	}
	return s
}

func (s *SubscriptionFields) NewErc721Transfers(argCriteria Optional[Erc721TransferFilterInput], selectFields func(*Erc721TransferFields)) *SubscriptionFields {
	child := s.node.object("newErc721Transfers", argument{name: "criteria", gqlType: "ERC721TransferFilterInput", value: argCriteria, optional: true})
	if selectFields != nil {
		selectFields(&Erc721TransferFields{node: child})
	}
	return s
}

func (s *SubscriptionFields) NewUniswapV2Swaps(argCriteria Optional[Erc20TokenSwapEventFilterInput], selectFields func(*Erc20TokenSwapEventFields)) *SubscriptionFields {
	child := s.node.object("newUniswapV2Swaps", argument{name: "criteria", gqlType: "ERC20TokenSwapEventFilterInput", value: argCriteria, optional: true})
	if selectFields != nil {
		selectFields(&Erc20TokenSwapEventFields{node: child})
	}
	return s
}

func (s *SubscriptionFields) NewDelegatedEvents(argCriteria Optional[DelegatedFilterInput], selectFields func(*DelegatedEventFields)) *SubscriptionFields {
	child := s.node.object("newDelegatedEvents", argument{name: "criteria", gqlType: "DelegatedFilterInput", value: argCriteria, optional: true})
	if selectFields != nil {
		selectFields(&DelegatedEventFields{node: child})
	}
	return s
}

func (s *SubscriptionFields) NewDelegateStakeChangedEvents(argCriteria Optional[DelegateStakeChangedFilterInput], selectFields func(*DelegateStakeChangedEventFields)) *SubscriptionFields {
	child := s.node.object("newDelegateStakeChangedEvents", argument{name: "criteria", gqlType: "DelegateStakeChangedFilterInput", value: argCriteria, optional: true})
	if selectFields != nil {
		selectFields(&DelegateStakeChangedEventFields{node: child})
	}
	return s
}

func (s *SubscriptionFields) NewPendingTransactions(selectFields func(*TransactionFields)) *SubscriptionFields {
	child := s.node.object("newPendingTransactions")
	if selectFields != nil {
		selectFields(&TransactionFields{node: child})
	}
	return s
}

func (s *SubscriptionFields) BlockReverted(selectFields func(*BlockRevertedEventFields)) *SubscriptionFields {
	child := s.node.object("blockReverted")
	if selectFields != nil {
		selectFields(&BlockRevertedEventFields{node: child})
	}
	return s
}
