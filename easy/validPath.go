package easy

//bfs
func validPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}
	info := map[int][]int{}
	for _, v := range edges {
		info[v[0]] = append(info[v[0]], v[1])
		info[v[1]] = append(info[v[1]], v[0])
	}
	accessed := map[int]struct{}{}
	work := make([]int, 0)

	accessed[source] = struct{}{}
	work = append(work, source)
	for len(work) > 0 {
		cur := work[0]
		for _, i := range info[cur] {
			if i == destination {
				return true
			}
			if _, exist := accessed[i]; !exist {
				work = append(work, i)
				accessed[i] = struct{}{}
			}
		}
		work = work[1:]
	}
	return false
}

func validPath_dfs(n int, edges [][]int, source int, destination int) bool {
	info := map[int][]int{}
	for _, v := range edges {
		info[v[0]] = append(info[v[0]], v[1])
		info[v[1]] = append(info[v[1]], v[0])
	}
	accessed := map[int]struct{}{}
	var dfs func(s, d int, accessed map[int]struct{}, info map[int][]int) bool
	dfs = func(s, d int, accessed map[int]struct{}, info map[int][]int) bool {
		if s == d {
			return true
		}
		if _, b := accessed[s]; b {
			return false
		}
		accessed[s]=struct{}{}
		for _, v := range info[s] {
			if dfs(v,d,accessed,info){
				return true
			}
		}
		return false
	}
	return dfs(source, destination, accessed, info)
}
