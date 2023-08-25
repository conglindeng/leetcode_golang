package middle

import (
	"bytes"
)

func PrintBin(num float64) string {
	b := bytes.Buffer{}
	b.WriteString("0.")
	for n := 0.5; num != 0.0; n = n / 2 {
		if num >= n {
			b.WriteString("1")
			num = num - n
		} else {
			b.WriteString("0")
		}
	}
	if b.Len() <= 32 {
		return b.String()
	}
	return "ERROR"
}
