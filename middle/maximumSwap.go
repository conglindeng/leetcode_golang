package middle

import (
	"strconv"
)

func maximumSwap(num int) int {
	res := num
	str := []byte(strconv.Itoa(num))
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			str[i], str[j] = str[j], str[i]
			m, _ := strconv.Atoi(string(str))
			res = max(res, m)
			str[i], str[j] = str[j], str[i]
		}
	}
	return res
}

func MaximumSwap_(num int) int {
	chars := []byte(strconv.Itoa(num))
	n := len(chars)
	maxInt, index1, index2 := n-1, 0, 0
	for i := n - 1; i >= 0; i-- {
		if chars[i] > chars[maxInt] {
			maxInt = i
		} else if chars[i] < chars[maxInt] {
			index1, index2 = maxInt, i
		}
	}
	if index2 == -1 {
		return num
	}
	chars[index1], chars[index2] = chars[index2], chars[index1]
	v, _ := strconv.Atoi(string(chars))
	return v
}
