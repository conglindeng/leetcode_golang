package easy

func minimumRecolors(blocks string, k int) int {
	wC := 0
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			wC++
		}
	}
	res := wC
	for i := k; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			wC++
		}
		if blocks[i-k] == 'W' {
			wC--
		}
		res = min(res, wC)
	}
	return res
}
