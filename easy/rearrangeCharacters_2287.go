package easy

import "math"

func rearrangeCharacters(s string, target string) int {
	need := map[rune]int{}
	for _, c := range target {
		need[c]++
	}
	count := map[rune]int{}
	for _, c := range s {
		count[c]++
	}
	res := math.MaxInt
	for k, v := range need {
		res = min(res, count[k]/v)
	}
	return res
}
