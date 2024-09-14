package middle

func removeStars(s string) string {
	var runes []rune
	for _, r := range s {
		if r != '*' {
			runes = append(runes, r)
		} else {
			runes = runes[:len(runes)-1]
		}
	}
	return string(runes)
}
