package middle

import (
	"bytes"
	"strings"
)

func RemoveComments(source []string) []string {
	// status: 0-common 1-// 2-/**/
	pre := 0
	var b bytes.Buffer
	for _, s := range source {
		for i := 0; i < len(s); i++ {
			c := s[i]
			if pre == 0 {
				if c == '/' {
					if i < len(s)-1 {
						if s[i+1] == '/' {
							pre = 1
							i++
						} else if s[i+1] == '*' {
							pre = 2
							i++
						} else {
							b.WriteByte(c)
						}
					} else {
						b.WriteByte(c)
					}
				} else {
					b.WriteByte(c)
				}
			} else if pre == 1 {
				continue
			} else if pre == 2 {
				if c == '*' {
					if i < len(s)-1 {
						if s[i+1] == '/' {
							pre = 0
							i++
						}
					}
				}
				continue
			}
		}
		if pre != 2 {
			b.WriteByte('\n')
		}
		if pre == 1 {
			pre = 0
		}
	}
	res := make([]string, 0)
	s := strings.Split(b.String(), "\n")
	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			res = append(res, s[i])
		}
	}
	return res
}
