package utils

// GetPointer returns a pointer to a value
func GetPointer[T any](p T) *T {
	return &p
}
