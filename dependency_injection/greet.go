package dependencyinjection

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "hello %s", name)
}
