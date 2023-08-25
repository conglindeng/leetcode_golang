package middle

func TotalFruit(fruits []int) int {
	left, right := 0, 0
	res := 0
	set := map[int]int{}
	for right < len(fruits) {
		cur := fruits[right]
		if len(set) < 2 {
			set[cur] += 1
			res = max(res, right-left+1)
		} else {
			if _, ok := set[cur]; ok {
				set[cur] += 1
				res = max(res, right-left+1)
			} else {
				set[cur] = 1
				for left < right {
					set[fruits[left]] -= 1
					if set[fruits[left]] == 0 {
						delete(set, fruits[left])
						left++
						break
					}
					left++
				}
				res = max(res, right-left+1)
			}
		}
		right++

	}
	return res
}

func totalFruit(fruits []int) int {
	count := map[int]int{}
	left, res := 0, 0
	for right, x := range fruits {
		count[x] += 1
		for len(count) > 2 {
			count[fruits[left]] -= 1
			if count[fruits[left]] == 0 {
				delete(count, fruits[left])
			}
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}
