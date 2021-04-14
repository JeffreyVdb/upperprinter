package printerz

import (
	"fmt"
	"io"
	"strings"
)

func FPrintUpper(w io.Writer, m string) (n int, err error) {
	stripped := strings.TrimSpace(m)
	upper := strings.ToUpper(stripped)
	return fmt.Fprintln(w, upper)
}
