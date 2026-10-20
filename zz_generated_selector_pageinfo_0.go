package hyperflow

func (s *PageInfoFields) Typename() *PageInfoFields {
	s.node.field("__typename")
	return s
}

func (s *PageInfoFields) HasNextPage() *PageInfoFields {
	s.node.field("hasNextPage")
	return s
}

func (s *PageInfoFields) HasPreviousPage() *PageInfoFields {
	s.node.field("hasPreviousPage")
	return s
}

func (s *PageInfoFields) TotalCount() *PageInfoFields {
	s.node.field("totalCount")
	return s
}

func (s *PageInfoFields) CurrentPage() *PageInfoFields {
	s.node.field("currentPage")
	return s
}

func (s *PageInfoFields) TotalPages() *PageInfoFields {
	s.node.field("totalPages")
	return s
}
