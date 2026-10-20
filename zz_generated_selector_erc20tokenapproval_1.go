package hyperflow

func (s *Erc20TokenApprovalFields) BlockTimestamp() *Erc20TokenApprovalFields {
	s.node.field("blockTimestamp")
	return s
}

func (s *Erc20TokenApprovalFields) ParsedApproval(selectFields func(*ParsedTokenAmountFields)) *Erc20TokenApprovalFields {
	child := s.node.object("parsedApproval")
	if selectFields != nil {
		selectFields(&ParsedTokenAmountFields{node: child})
	}
	return s
}
