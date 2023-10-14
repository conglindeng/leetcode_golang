package middle

import (
	"sort"
	"strings"
)

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	scores := make([][2]int, 0)
	positiveMap, negativeMap := map[string]bool{}, map[string]bool{}
	for _, v := range positive_feedback {
		positiveMap[v] = true
	}
	for _, v := range negative_feedback {
		negativeMap[v] = true
	}
	for i := 0; i < len(report); i++ {
		id := student_id[i]
		positive, negative := 0, 0
		for _, s := range strings.Split(report[i], " ") {
			if negativeMap[s] {
				negative++
			} else if positiveMap[s] {
				positive++
			}
		}
		scores = append(scores, [2]int{id, positive*3 - negative})
	}
	sort.Slice(scores, func(i, j int) bool {
		if scores[i][1] == scores[j][1] {
			return scores[i][0] < scores[j][0]
		}
		return scores[i][1] < scores[j][1]
	})
	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, scores[i][0])
	}
	return res
}
