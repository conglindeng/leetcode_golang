package easy

import "bytes"

func bestHand(ranks []int, suits []byte) string {
	if bytes.Count(suits, suits[:1]) == 5 {
		return "Flush"
	}
	rCount, pair := map[int]int{}, false
	for _, n := range ranks {
		rCount[n]++
		if rCount[n] >= 3 {
			return "Three of a Kind"
		} else if rCount[n] == 2 {
			pair = true
		}
	}
	if pair {
		return "Pair"
	}
	return "High Card"
}
