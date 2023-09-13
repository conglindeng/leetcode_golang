package middle

//todo: dcl check it again 
func CheckIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	edges := make([][]int, numCourses)
	inDegree := map[int]int{}
	isPre := make([][]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		isPre[i] = make([]bool, numCourses)
		edges[i] = make([]int, 0)
	}
	for _, v := range prerequisites {
		inDegree[v[1]]++
		edges[v[0]] = append(edges[v[0]], v[1])
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
		for _, v := range edges[cur] {
			isPre[cur][v] = true
			for i := 0; i < numCourses; i++ {
				isPre[i][v] = isPre[i][v] || isPre[i][cur]
			}
			inDegree[v]--
			if inDegree[v] == 0 {
				stack = append(stack, v)
			}
		}
	}
	res := make([]bool, len(queries))
	for i, v := range queries {
		res[i] = isPre[v[0]][v[1]]
	}
	return res
}

func checkIfPrerequisite_dfs(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	isPre := make([][]bool, numCourses)
	edges := make([][]int, numCourses)
	visited := make([]bool, numCourses)
	var dfs func(cur int)
	dfs = func(cur int) {
		if visited[cur] {
			return
		}
		visited[cur] = true
		for _, v := range edges[cur] {
			isPre[cur][v] = true
			dfs(v)
			for i := 0; i < numCourses; i++ {
				isPre[cur][i] = isPre[cur][i] || isPre[v][i]
			}
		}
	}
	for i := 0; i < numCourses; i++ {
		edges[i] = []int{}
		isPre[i] = make([]bool, numCourses)
	}
	for _, v := range prerequisites {
		edges[v[0]] = append(edges[v[0]], v[1])
	}
	for i := 0; i < numCourses; i++ {
		dfs(i)
	}
	res := []bool{}
	for _, v := range queries {
		res = append(res, isPre[v[0]][v[1]])
	}
	return res
}
