package easy

import (
	"fmt"
	"sort"
)

func MaxmiumScore(cards []int, cnt int) int {
	if len(cards) < cnt {
		return 0
	}
	sort.Ints(cards)
	for _, i := range cards {
		fmt.Print(i)
		fmt.Print(",")
	}

	sum := 0
	for i := 1; i < cnt; i++ {
		sum += cards[len(cards)-i]
	}
	fmt.Println(sum)
	fmt.Println(len(cards) - cnt)
	for i := len(cards) - cnt; i >= 0; i-- {
		if (sum+cards[i])%2 == 0 {
			return sum + cards[i]
		}
	}
	return 0
}
