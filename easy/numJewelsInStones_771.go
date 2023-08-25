package easy

func numJewelsInStones(jewels string, stones string) int {
	jewel := map[rune]struct{}{}
	for _, r := range jewels {
		jewel[r] = struct{}{}
	}
	res := 0
	for _, s := range stones {
		if _, ok := jewel[s]; ok {
			res++
		}
	}
	return res
}
