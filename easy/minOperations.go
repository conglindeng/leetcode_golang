package easy

func minOperations(logs []string) int {
	deep := 0
	for _, log := range logs {
		if log == "../" {
            if deep>0{
			    deep--
            }
		} else if log != "./" {
			deep++
		}
	}
	return deep
}
