package easy

func CheckIfPangram(sentence string) bool {
	l := len(sentence)
	if l < 26 {
		return false
	}
	cnt := [26]int{}
	for i := 0; i < l; i++ {
		cnt[sentence[i]-'a']++
	}
	for _, i := range cnt {
		if i == 0 {
			return false
		}
	}
	return true
}
