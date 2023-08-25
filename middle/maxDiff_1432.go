package middle

import (
	"strconv"
	"strings"
)

func MaxDiff(num int) int {
	s := strconv.Itoa(num)
	ma := 0
	tmp := s
	old, new := byte('-'), byte('-')
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '9' {
			new = '9'
			old = c
			break
		}
	}
	ma = replaceOldbyNew(tmp, new, old)
	tmp = s
	mi := 0
	new = '-'
	var firstByte byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c > '0' {
			if i == 0 {
				firstByte = c
				if c == '1' {
					continue
				}
				new = '1'
			} else {
				if c == firstByte {
					continue
				}
				new = '0'
			}
			old = c
			break
		}
	}
	mi = replaceOldbyNew(tmp, new, old)
	return ma - mi
}

func replaceOldbyNew(tmp string, new, old byte) int {
	if new == '-' {
		n, _ := strconv.Atoi(tmp)
		return n
	}
	n, _ := strconv.Atoi(strings.ReplaceAll(tmp, string(old), string(new)))
	return n
}
