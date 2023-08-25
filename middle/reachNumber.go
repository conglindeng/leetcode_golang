package middle

import "math"

func ReachNumber(target int) int {
	res := math.MaxInt
	var dfs func(position, step, target int)
	dfs = func(position, step, target int) {
		if step == target*target {
			return
		}
		if position == target {
			res = min(res, step)
		} else if position > target {
			return
		} else {
			dfs(position+step, step+1, target)
			dfs(position-step, step+1, target)
		}
	}
	dfs(0, 1, target)
	return res
}
