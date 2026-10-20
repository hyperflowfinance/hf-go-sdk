package hyperflow

func (s *EffectivenessFields) Typename() *EffectivenessFields {
	s.node.field("__typename")
	return s
}

func (s *EffectivenessFields) Effectiveness() *EffectivenessFields {
	s.node.field("effectiveness")
	return s
}

func (s *EffectivenessFields) JsonrpcRatio() *EffectivenessFields {
	s.node.field("jsonrpcRatio")
	return s
}
