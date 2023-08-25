package easy

func mostFrequentEven(nums []int) int {
	count := map[int]int{}
	res := -1
	maxC := 0
	for _, n := range nums {
		if n%2 == 0 {
			count[n]++
			if count[n] > maxC {
				res = n
				maxC = count[n]
			} else if count[n] == maxC {
				res = min(res, n)
			}
		}
	}
	return res
}
