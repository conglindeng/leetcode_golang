package middle

import "strconv"

func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := "0"
	for i := len(num1) - 1; i >= 0; i-- {
		multiRes := multiSigleNum(num1[i], num2) + completeZero(len(num1)-1-i)
		res = addTwoStr(multiRes, res)
	}
	return res
}

func addTwoStr(num1, num2 string) string {
	carry := 0
	i, j := len(num1)-1, len(num2)-1
	res := []int{}
	for i >= 0 && j >= 0 {
		first, second := int(num1[i]-'0'), int(num2[j]-'0')
		cur := first + second + carry
		res = append(res, cur%10)
		carry = cur / 10
		i--
		j--
	}
	for i >= 0 {
		cur := int(num1[i]-'0') + carry
		res = append(res, cur%10)
		carry = cur / 10
		i--
	}
	for j >= 0 {
		cur := int(num2[j]-'0') + carry
		res = append(res, cur%10)
		carry = cur / 10
		j--
	}
	if carry > 0 {
		res = append(res, carry)
	}
	return convertNums2Str(res)
}

func multiSigleNum(baseStr byte, multipier string) string {
	res := "0"
	base := int(baseStr - '0')
	for i := len(multipier) - 1; i >= 0; i-- {
		cur := strconv.Itoa(int(multipier[i]-'0')*base) + completeZero(len(multipier)-1-i)
		res = addTwoStr(res, cur)
	}
	return res
}

func convertNums2Str(nums []int) string {
	res := ""
	for i := 0; i < len(nums); i++ {
		res = strconv.Itoa(nums[i]) + res
	}
	return res
}

func completeZero(n int) string {
	res := ""
	for i := 0; i < n; i++ {
		res += "0"
	}
	return res
}