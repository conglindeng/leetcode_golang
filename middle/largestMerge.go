package middle

func largestMerge(word1 string, word2 string) string {
	merge := []byte{}
	i, j, m, n := 0, 0, len(word1), len(word2)
	for i < m || j < n {
		if i < m && word1[i:] > word2[j:] {
			merge = append(merge, word1[i])
			i++
		} else {
			merge = append(merge, word2[j])
			j++
		}
	}
	return string(merge)
}
