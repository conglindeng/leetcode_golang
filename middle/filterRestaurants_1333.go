package middle

import "sort"

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	tmp := make([][]int, 0)
	for _, v := range restaurants {
		if veganFriendly == 1 && v[2] != veganFriendly {
			continue
		}
		if v[3] > maxPrice {
			continue
		}
		if v[4] > maxDistance {
			continue
		}
		tmp = append(tmp, []int{v[0], v[1]})
	}
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i][1] == tmp[j][1] {
			return tmp[i][0] < tmp[j][0]
		}
		return tmp[i][1] > tmp[j][1]
	})
	res := []int{}

	for _, v := range tmp {
		res = append(res, v[0])
	}
	return res
}
