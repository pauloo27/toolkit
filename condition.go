package toolkit

// Conditional[T any] works like a ternary operator. If condition is true,
// then ifTrue is returned, otherwise ifFalse is returned.
func Conditional[T any](condition bool, ifTrue T, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}
