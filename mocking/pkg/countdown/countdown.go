package countdown

import (
	"fmt"
	"io"
	"time"
)

func Countdown(writer io.Writer) {
	const startFrom int = 3
	for i := startFrom; i > 0; i-- {
		fmt.Fprintf(writer, "%d\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(writer, "Go!\n")
}
