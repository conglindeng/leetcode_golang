package easy

import (
	"fmt"
	"reflect"
)

func digitCount(num string) bool {
	require := map[int]int{}
	count := map[int]int{}
	for i, c := range num {
		n := int(c - '0')
		count[n]++
		if n != 0 {
			require[i] = n
		}
	}
	fmt.Println()
	return reflect.DeepEqual(require, count)
}
