package easy

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	count1, count2 := [26]int{}, [26]int{}
	for _, c := range s1 {
		count1[c-'a']++
	}
	for _, c := range s2 {
		count2[c-'a']++
	}
	// for i := 0; i < 26; i++ {
	// 	if count1[i] != count2[i] {
	// 		return false
	// 	}
	// }
	return count1 == count2
}
