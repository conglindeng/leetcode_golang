package difficult

func FindItinerary(tickets [][]string) []string {
	edges := map[string]map[string]struct{}{}
	for _, v := range tickets {
		edges[v[0]] = map[string]struct{}{}
		edges[v[0]][v[1]] = struct{}{}
	}
	resCollection := make([][]string, 0)
	res := []string{}
	var find func(cur string)
	find = func(cur string) {
		res = append(res, cur)
		nexts := edges[cur]
		if len(nexts) == 0 {
			if len(res) == len(tickets)+1 {
				resCollection = append(resCollection, res)
			}
			return
		}
		for k, v := range nexts {
			delete(nexts, k)
			find(k)
			nexts[k] = v
		}

		// for j := 0; j < len(nexts); j++ {
		// 	next := nexts[j]
		// 	nexts[j] = nexts[len(nexts)-1]
		// 	edges[cur] = nexts[:len(nexts)-1]
		// }
		res = res[:len(res)-1]
	}
	find("JFK")
	// sort.Slice(resCollection,func(i, j int) bool {
	// 	first,second:=resCollection[i],resCollection[j]
	// })

		//todo: dcl

	return resCollection[0]
}
