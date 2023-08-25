package easy

func greatestLetter(s string) string {
	exist := map[byte]struct{}{}
	var res byte
	for i := 0; i < len(s); i++ {
		c := byte(s[i])
		exist[c] = struct{}{}
		if c <= 'Z' {
			if _, b := exist[c+('a'-'A')]; b {
				if res < c {
					res = c
				}
			}
		} else {
			if _, b := exist[c-('a'-'A')]; b {
				if res < c-('a'-'A') {
					res = c - ('a' - 'A')
				}
			}
		}
	}
	if res == 0 {
		return ""
	}
	return string(res)
}
