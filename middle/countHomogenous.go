package middle

func countHomogenous(s string) int {
	const mod int = 1e9 + 7
	left, right := 0, 0
	res := 0
	for left < len(s) {
		for right < len(s) && s[right] == s[left] {
			right++
		}
		n := right - left
		res = res%mod + (n*(n+1)/2)%mod
		left = right
	}
	return res
}
