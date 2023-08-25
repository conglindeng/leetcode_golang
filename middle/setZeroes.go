package middle

import "fmt"

func SetZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	// res := make([][]int, m)
	// for i := range res {
	// 	res[i] = make([]int, n)
	// }
	row := make([]bool, m)
	col := make([]bool, n)
	for i, nums := range matrix {
		for j, num := range nums {
			if num == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i, nums := range matrix {
		for j, _ := range nums {
			if col[i] || row[j] {
				matrix[i][j] = 0
			}
		}
	}
	fmt.Print(matrix)
}
