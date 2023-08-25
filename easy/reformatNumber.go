package easy

import (
	"bytes"
	"strings"
)

func ReformatNumber(number string) string {
	number = strings.ReplaceAll(number, " ", "")
	number = strings.ReplaceAll(number, "-", "")
	var b bytes.Buffer
	l := len(number)
	m := 0
	for i, c := range number {
		if m%3 == 0 && l-i <= 4 {
			if b.Len() > 0 {
				b.WriteByte('-')
			}
			if l-i == 4 {
				b.WriteString(number[i : i+2])
				b.WriteByte('-')
				b.WriteString(number[i+2:])
			} else {
				b.WriteString(number[i:l])
			}
			break
		} else {
			if i != 0 && i%3 == 0 {
				b.WriteByte('-')
			}
			b.WriteRune(c)
			m++
		}
	}
	return b.String()
}

func ReformatNumber_(number string) string {
	number = strings.ReplaceAll(number, " ", "")
	number = strings.ReplaceAll(number, "-", "")
	var b bytes.Buffer
	l := len(number)
	if l <= 3 {
		return number
	}
	m := l / 3
	n := l % 3
	i := 0
	for ; i < m-1; i++ {
		if i != 0 {
			b.WriteByte('-')
		}
		b.WriteString(number[i*3 : (i+1)*3])
	}
	if b.Len() > 0 {
		b.WriteByte('-')
	}
	if n == 1 {
		b.WriteString(number[(i)*3 : (i+1)*3-1])
		b.WriteByte('-')
		b.WriteString(number[(i+1)*3-1:])

	} else if n == 0 {
		b.WriteString(number[(i)*3:])
	} else if n == 2 {
		b.WriteString(number[(i)*3 : (i+1)*3])
		b.WriteByte('-')
		b.WriteString(number[(i+1)*3:])
	}
	return b.String()
}
