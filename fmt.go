package toolkit

import "fmt"

// Fmtf formats according to a format specifier and returns the resulting string.
// It just call fmt.Sprintf, nothing special
func Fmt(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
