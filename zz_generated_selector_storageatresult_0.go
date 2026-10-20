package hyperflow

func (s *StorageAtResultFields) Typename() *StorageAtResultFields {
	s.node.field("__typename")
	return s
}

func (s *StorageAtResultFields) Value() *StorageAtResultFields {
	s.node.field("value")
	return s
}
