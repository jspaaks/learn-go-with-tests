package iteration

import "strings"

const nRepeat int = 5

func Repeat(s string) string {	
	var repeated strings.Builder
	for i := 0; i < nRepeat; i++ {
		repeated.WriteString(s)
	}
	return repeated.String()
}
