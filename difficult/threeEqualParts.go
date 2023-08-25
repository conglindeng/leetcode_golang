package difficult

func threeEqualParts(arr []int) []int {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	if sum%3 != 0 {
		return []int{-1, -1}
	}
	if sum == 0 {
		return []int{0, 2}
	}
	first, second, third, cur := 0, 0, 0, 0
	partial := sum / 3
	for i, n := range arr {
		if n == 1 {
			if cur == 0 {
				first = i
			} else if cur == partial {
				second = i
			} else if cur == 2*partial {
				third = i
			}
			cur++
		}
	}
	l := len(arr) - third
	if first+l <= second && second+l <= third {
		i := 0
		for third+i < len(arr) {
			if arr[first+i] != arr[second+i] || arr[first+i] != arr[third+i] {
				return []int{-1, -1}
			}
			i++
		}
		return []int{first + l - 1, second + l}
	}
	return []int{-1, -1}
}
