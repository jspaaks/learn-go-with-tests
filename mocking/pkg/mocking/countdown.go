package countdown

import (
	"bytes"
	"fmt"
)

func Countdown(buffer *bytes.Buffer) {
	const startFrom int = 3
	for i := startFrom; i > 0; i-- {
		fmt.Fprintf(buffer, "%d\n", i)
	}
	fmt.Fprint(buffer, "Go!\n")
}
