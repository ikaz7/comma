// Comma inserts commas in non-negative decimal integer strings non-recursively
package comma

import (
	"bytes"
	"fmt"
	"os"
	"unicode"
)

func Parse(s []byte) string {
	var f []byte // holds fractional part
	d := bytes.Index(s, []byte("."))
	if d >= 0 {
		f = []byte(s[d:])
		s = s[:d]
	}
	var buf bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		if !unicode.IsNumber(rune(s[len(s)-(i+1)])) {
			fmt.Fprintf(os.Stderr, "comma: %s is not a number\n", s)
			os.Exit(1)
		}
		buf.WriteByte(s[len(s)-(i+1)])
		if i != 0 && i%3 == 0 {
			buf.WriteByte(',')
		}
	}
	if d >= 0 {
		buf.Write(f)
	}
	return buf.String()
}
