package hyperflow

func (s *MutationFields) Typename() *MutationFields {
	s.node.field("__typename")
	return s
}

func (s *MutationFields) SendRawTransaction(argSignedTx string, argChain Optional[string], argChainID Optional[int32]) *MutationFields {
	s.node.field("sendRawTransaction", argument{name: "signedTx", gqlType: "String!", value: argSignedTx, optional: false}, argument{name: "chain", gqlType: "String", value: argChain, optional: true}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	return s
}

func (s *MutationFields) SendTransaction(argTxData BuildTransactionInput, argSignature SignatureInput, argChain Optional[string], argChainID Optional[int32]) *MutationFields {
	s.node.field("sendTransaction", argument{name: "txData", gqlType: "BuildTransactionInput!", value: argTxData, optional: false}, argument{name: "signature", gqlType: "SignatureInput!", value: argSignature, optional: false}, argument{name: "chain", gqlType: "String", value: argChain, optional: true}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	return s
}

func (s *MutationFields) SendTransactions(argNode TransactionNodeInput, argChain Optional[string], argChainID Optional[int32], selectFields func(*TransactionResultFields)) *MutationFields {
	child := s.node.object("sendTransactions", argument{name: "node", gqlType: "TransactionNode!", value: argNode, optional: false}, argument{name: "chain", gqlType: "String", value: argChain, optional: true}, argument{name: "chainId", gqlType: "Int", value: argChainID, optional: true})
	if selectFields != nil {
		selectFields(&TransactionResultFields{node: child})
	}
	return s
}

func (s *MutationFields) SignTransaction(argTxData BuildTransactionInput, argSignature SignatureInput) *MutationFields {
	s.node.field("signTransaction", argument{name: "txData", gqlType: "BuildTransactionInput!", value: argTxData, optional: false}, argument{name: "signature", gqlType: "SignatureInput!", value: argSignature, optional: false})
	return s
}
