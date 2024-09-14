package easy

func removeTrailingZeros(num string) string {
	j := len(num) - 1
	for ; j >= 0 && num[j] == '0'; j-- {

	}
	return num[:j+1]
}
