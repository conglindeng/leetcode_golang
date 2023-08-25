package easy

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	count := map[string]map[string]int{}
	count["type"] = map[string]int{}
	count["color"] = map[string]int{}
	count["name"] = map[string]int{}
	for _, item := range items {
		count["type"][item[0]]++
		count["color"][item[1]]++
		count["name"][item[2]]++
	}
	return count[ruleKey][ruleValue]
}

func countMatches_(items [][]string, ruleKey string, ruleValue string) int {
	idx := map[string]int{"type": 0, "color": 1, "name": 2}
	count := 0
	for _, item := range items {
		if item[idx[ruleKey]] == ruleValue {
			count++
		}
	}
	return count
}
