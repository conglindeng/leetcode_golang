package middle

//为了防止重复计算，我们可以设 \textit{arr}[i]arr[i] 左边的元素都必须满足小于等于 \textit{arr}[i]arr[i]，\textit{arr}[i]arr[i] 右边的元素必须满足严格小于 \textit{arr}[i]arr[i]。
//???  not understand
func SumSubarrayMins(arr []int) int {
	const mod int = 1e9 + 7
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)
	stack := []int{}
	for i, x := range arr {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= x {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			//why it is i+1 , but not i
			left[i] = i + 1
		} else {
			left[i] = i - stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[i] < arr[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n - i
		} else {
			right[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	res := 0
	for i, l := range left {
		res = (res + arr[i]*l*right[i]) % mod
	}
	return res
}
