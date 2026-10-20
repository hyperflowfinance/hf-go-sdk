package hyperflow

func (s *EstimateGasResultFields) Typename() *EstimateGasResultFields {
	s.node.field("__typename")
	return s
}

func (s *EstimateGasResultFields) Gas() *EstimateGasResultFields {
	s.node.field("gas")
	return s
}

func (s *EstimateGasResultFields) GasDecimal() *EstimateGasResultFields {
	s.node.field("gasDecimal")
	return s
}
