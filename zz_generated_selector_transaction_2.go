package hyperflow

func (s *TransactionFields) InternalTransactions(selectFields func(*InternalTransactionFields)) *TransactionFields {
	child := s.node.object("internalTransactions")
	if selectFields != nil {
		selectFields(&InternalTransactionFields{node: child})
	}
	return s
}
