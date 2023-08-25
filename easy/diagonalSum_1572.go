package easy

func diagonalSum(mat [][]int) int {
	n := len(mat)
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				res += mat[i][j]
			} else if i+j == n-1 {
				res += mat[i][j]
			}
		}
	}
	return res
}
