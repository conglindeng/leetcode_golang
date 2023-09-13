package middle

func CanFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges  = make([][]int, numCourses)
		visit  = map[int]int{}
		valid  = true
		result = []int{}
		dfs    func(v int)
	)
	dfs = func(v int) {
		visit[v] = 1
		for _, n := range edges[v] {
			if visit[n] != 1 {
				dfs(n)
				if !valid {
					return
				}
			} else if visit[n] == 1 {
				valid = false
				return
			}
		}
		visit[v] = 2
		result = append(result, v)
	}
	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visit[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

func CanFinish_bfs(numCourses int, prerequisites [][]int) bool {
	inDegree := map[int]int{}
	preCondition := map[int][]int{}
	stack := []int{}
	for i := 0; i < numCourses; i++ {
		inDegree[i] = 0
	}
	for _, v := range prerequisites {
		preCondition[v[0]] = append(preCondition[v[0]], v[1])
		inDegree[v[1]]++
	}
	for k, v := range inDegree {
		if v == 0 {
			stack = append(stack, k)
		}
	}
	for len(stack) > 0 {
		for len(stack) > 0 {
			cur := stack[0]
			stack = stack[1:]
			for _, v := range preCondition[cur] {
				inDegree[v]--
			}
			inDegree[cur] = -1
		}
		for k, v := range inDegree {
			if v == 0 {
				stack = append(stack, k)
			}
		}
	}
	for _, v := range inDegree {
		if v > 0 {
			return false
		}
	}
	return true
}

func CanFinish_bfs_(numCourses int, prerequisites [][]int) bool {
	edges := map[int][]int{}
	inDegree := map[int]int{}
	res := []int{}
	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
		inDegree[v[0]]++
	}
	stack := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			stack = append(stack, i)
		}
	}
	for len(stack) > 0 {
		cur := stack[0]
		res = append(res, cur)
		stack = stack[1:]
		for _, m := range edges[cur] {
			inDegree[m]--
			if inDegree[m] == 0 {
				stack = append(stack, m)
			}
		}
	}
	return len(res) == numCourses
}
