package middle

func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := map[int][]int{}
	inDegree := map[int]int{}
	res := []int{}

	for _, v := range prerequisites {
		inDegree[v[0]]++
		edges[v[1]] = append(edges[v[1]], v[0])
	}
	stack := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			stack = append(stack, i)
		}
	}
	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]
		res = append(res, cur)
		for _, v := range edges[cur] {
			inDegree[v]--
			if inDegree[v] == 0 {
				stack = append(stack, v)
			}
		}
	}
	if len(res) == numCourses {
		return res
	}
	return nil
}
