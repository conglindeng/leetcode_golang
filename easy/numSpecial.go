package easy

func numSpecial(mat [][]int) int {
	row := make([]int, len(mat))
	col := make([]int, len(mat[0]))
	for i, r := range mat {
		for j, c := range r {
			row[i] += c
			col[j] += c
		}
	}
	res := 0
	for i, r := range mat {
		for j, c := range r {
			if c == 1 && row[i] == 1 && col[j] == 1 {
				res++
			}
		}
	}
	return res
}
