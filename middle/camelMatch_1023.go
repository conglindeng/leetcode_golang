package middle

import "unicode"

func CamelMatch(queries []string, pattern string) []bool {
	res := []bool{}
	for _, s := range queries {
		res = append(res, doCamelMatch(s, pattern))
	}
	return res
}
func doCamelMatch(query, pattern string) bool {
	first := getUpperLetterCount(query)
	second := getUpperLetterCount(pattern)
	if len(first) != len(second) {
		return false
	}
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func getUpperLetterCount(s string) []byte {
	res := []byte{}
	for _, c := range s {
		if unicode.IsUpper(c) {
			res = append(res, byte(c))
		}
	}
	return res
}
