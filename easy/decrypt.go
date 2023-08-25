package easy

func decrypt(code []int, k int) []int {
	//traverse back from the current node
	direct := 1
	step := k
	if k < 0 {
		//traverse forward from the current node
		direct = -1
		step = -k
	}
	res := make([]int, len(code))
	for i := 0; i < len(code); i++ {
		sum := 0
		for m := 1; m <= step; m++ {
			sum += code[(i+direct*m+len(code))%len(code)]
		}
		res[i] = sum
	}
	return res
}

func Decrypt_(code []int, k int) []int {
	res := make([]int, len(code))
	if k == 0 {
		return res
	}
	step := k
	if k < 0 {
		step = -k
	}
	//calculate sum by sliding window , then find the sum`s position.
	//by the way,assume the window range is [left,right],if k is positive ,
	//the position is [left-1+len(code)%len(code)],otherwise is [right+1+len(code)%len(code)]
	left, right, sum := 0, 0, 0
	for left < len(code) {
		for right-left < step {
			sum += code[right%len(code)]
			right++
		}
		if k < 0 {
			res[(right+len(code))%len(code)] = sum
		} else {
			res[(left-1+len(code))%len(code)] = sum
		}
		sum -= code[left]
		left++
	}
	return res
}
