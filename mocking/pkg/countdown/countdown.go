package countdown

import (
	"fmt"
	"io"
)

func Countdown(writer io.Writer, sleeper Sleeper) {
	const startFrom int = 3
	for i := startFrom; i > 0; i-- {
		fmt.Fprintf(writer, "%d\n", i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, "Go!\n")
}
