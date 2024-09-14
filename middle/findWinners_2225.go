package middle

import "sort"

func findWinners(matches [][]int) [][]int {
	loseCount := map[int]int{}
	for _, match := range matches {
		loseCount[match[1]]++
		loseCount[match[0]] += 0
	}
	bigWinner := []int{}
	oneTime := []int{}
	for k, v := range loseCount {
		if v == 0 {
			bigWinner = append(bigWinner, k)
		}
		if v == 1 {
			oneTime = append(oneTime, k)
		}
	}
	sort.Ints(bigWinner)
	sort.Ints(oneTime)
	return [][]int{bigWinner, oneTime}
}
