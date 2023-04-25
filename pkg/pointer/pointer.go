package pointer

// String returns a pointer to a string containing the given value.
func String(s string) *string {
	return &s
}

func Bool(b bool) *bool {
	return &b
}
