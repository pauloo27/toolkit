package toolkit

// Is[T any] works like a ternary operator. If condition is true,
// then ifTrue is returned, otherwise ifFalse is returned.
func Is[T any](condition bool, ifTrue T, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}
