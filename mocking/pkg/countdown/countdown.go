package countdown

import (
	"fmt"
	"io"
)

func Countdown(writer io.Writer) {
	const startFrom int = 3
	for i := startFrom; i > 0; i-- {
		fmt.Fprintf(writer, "%d\n", i)
	}
	fmt.Fprint(writer, "Go!\n")
}
