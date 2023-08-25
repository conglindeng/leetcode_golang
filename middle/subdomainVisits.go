package middle

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	count := map[string]int{}
	for _, s := range cpdomains {
		strs := strings.Split(s, " ")
		c, _ := strconv.Atoi(strs[0])
		count[strs[1]] += c
		for i := 0; i < len(strs[1]); i++ {
			if strs[1][i] == '.' {
				count[strs[1][i+1:]] += c
			}
		}
	}
	res := make([]string, 0)
	for k, v := range count {
		res = append(res, fmt.Sprintf("%d %s", v, k))
	}
	return res
}
