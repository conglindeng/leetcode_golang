package middle

import "math"

func repairCars(ranks []int, cars int) int64 {
	l, r := 1, ranks[0]*cars*cars
	check := func(time int) bool {
		repaired := 0
		for _, v := range ranks {
			repaired += int(math.Sqrt(float64(time / v)))
		}
		return repaired >= cars
	}
	for l < r {
		mid := (r-l)/2 + l
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return int64(l)
}
