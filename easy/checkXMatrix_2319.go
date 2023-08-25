package easy

func checkXMatrix(grid [][]int) bool {
	l := len(grid)
	for i, nums := range grid {
		for j, n := range nums {
			if i == j || i+j == l-1 {
				if n == 0 {
					return false
				}
			} else {
				if n != 0 {
					return false
				}
			}
		}
	}
	return true
}
