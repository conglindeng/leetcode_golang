package easy

import (
	"strconv"
	"strings"
)

func AddStrings(num1 string, num2 string) string {
	carry := 0
	l1 := len(num1) - 1
	l2 := len(num2) - 1
	res := []string{}
	for {
		v := int(num1[l1]-'0') + int(num2[l2]-'0') + carry
		res = append(res, strconv.Itoa(v%10))
		carry = v / 10
		l1--
		l2--
		if l1 < 0 || l2 < 0 {
			break
		}
	}
	for l1 >= 0 {
		v := int(num1[l1]-'0') + carry
		res = append(res, strconv.Itoa(v%10))
		carry = v / 10
		l1--
	}
	for l2 >= 0 {
		v := int(num2[l2]-'0') + carry
		res = append(res, strconv.Itoa(v%10))
		carry = v / 10
		l2--
	}
	if carry > 0 {
		res = append(res, strconv.Itoa(carry))
	}
	//reverse
	for m, k := 0, len(res)-1; m < k; m, k = m+1, k-1 {
		res[m], res[k] = res[k], res[m]
	}
	return strings.Join(res, "")
}
