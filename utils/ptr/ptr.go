package ptr

func ToPtr[T any](v T) *T {
	return &v
}

func StringNotNil(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}
