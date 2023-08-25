package middle

import (
	"bytes"
	"strings"
	"unicode"
)

func maskPII(s string) string {
	if unicode.IsLetter(rune(s[0])) {
		return maskMail(s)
	} else {
		return maskPhone(s)
	}
}

func maskMail(s string) string {
	var b bytes.Buffer
	step := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if i == 0 {
			b.WriteRune(unicode.ToLower(rune(c)))
		} else if step == 1 {
			b.WriteRune(unicode.ToLower(rune(c)))
		} else if s[i+1] == '@' {
			b.WriteString("*****")
			b.WriteRune(unicode.ToLower(rune(c)))
			b.WriteByte('@')
			step++
			i++
		}
	}

	return b.String()
}

func maskPhone(s string) string {
	var b bytes.Buffer
	l := 0
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			l++
		}
	}
	for i := 0; i < l-10; i++ {
		b.WriteByte('*')
	}
	if b.Len() > 0 {
		b.WriteByte('-')
	}
	b.WriteString("***-***-")
	count := 0
	nums := make([]byte, 4)
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[len(s)-1-i])) {
			nums[3-count] = s[len(s)-1-i]
			count++
		}
		if count == 4 {
			break
		}
	}
	b.WriteString(string(nums))
	return b.String()
}

func maskPII_(s string) string {
	idx := strings.Index(s, "@")
	if idx > 0 {
		tmp := strings.ToLower(s)
		return string(string(tmp[0]) + "*****" + tmp[idx-1:])
	}
	var b bytes.Buffer
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			b.WriteByte(s[i])
		}
	}
	prefix := []string{"", "+*-", "+**-", "+***-"}
	s = b.String()
	return string(prefix[len(s)-10] + "***-***-" + s[len(s)-4:])
}
