package middle

import "unicode"

func findLongestSubarray(array []string) []string {
	// if letter add 1, else delete 1
	count := 0
	countInfo := map[int]int{}
	maxLength, startIdx := -1, -1
	countInfo[0] = -1
	for i, m := range array {
		if unicode.IsDigit(rune(m[0])) {
			count--
		} else {
			count++
		}
		if _, ok := countInfo[count]; ok {
			if i-countInfo[count] > maxLength {
				maxLength = i - countInfo[count]
				startIdx = countInfo[count]
			}
		} else {
			countInfo[count] = i
		}
	}
	if maxLength < 0 {
		return nil
	}
	return array[startIdx : startIdx+maxLength]
}
