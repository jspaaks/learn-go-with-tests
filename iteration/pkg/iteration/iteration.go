package iteration

import "strings"

// Returns its input argument s, repeated n times.
func Repeat(s string, n int) (repeated string) {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(s)
	}
	repeated = builder.String()
	return
}
