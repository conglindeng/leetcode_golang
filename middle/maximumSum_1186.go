package middle

func maximumSum(arr []int) int {
	//dp0 -> pick up all elements
	//dp1 -> drop one element
	dp0, dp1, res := arr[0], 0, arr[0]
	for i := 1; i < len(arr); i++ {
		n := arr[i]
		dp0, dp1 = max(dp0, 0)+n, max(dp1+n, dp0)
		res = max(res, max(dp0, dp1))
	}
	return res
}
