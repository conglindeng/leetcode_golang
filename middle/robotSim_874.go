package middle

func RobotSim(commands []int, obstacles [][]int) int {
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	direIdx := 0
	position := []int{0, 0}
	res := 0
	for _, v := range commands {
		if v == -1 {
			direIdx += 1
		} else if v == -2 {
			direIdx += 3
		} else {
			direction := directions[direIdx%4]
			position = robotMeetStoneOrNot(position, direction, v, obstacles)
			res = max(res, position[0]*position[0]+position[1]*position[1])
		}
	}
	return res
}

func robotMeetStoneOrNot(position, direction []int, step int, obstacles [][]int) []int {
	for _, v := range obstacles {
		//maybe meet
		needGo := []int{v[0] - position[0], v[1] - position[1]}
		if needGo[0] == 0 {
			if haveSameSign(needGo[1], direction[1]) && intAbs(v[1], position[1]) < step {
				return []int{v[0] - direction[0], v[1] - direction[1]}
			}
		}
		if needGo[1] == 0 {
			if haveSameSign(needGo[0], direction[0]) && intAbs(v[0], position[0]) < step {
				return []int{v[0] - direction[0], v[1] - direction[1]}
			}
		}
	}
	return []int{position[0] + step*direction[0], position[1] + direction[1]*step}
}

func haveSameSign(a, b int) bool {
	return (a > 0 && b > 0) || (a < 0 && b < 0)
}

// [4,-1,4,-2,4]
// [[2,4]]
func RobotSim_hash(commands []int, obstacles [][]int) int {
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	direIdx := 0
	px, py := 0, 0
	exist := map[int]bool{}
	for _, v := range obstacles {
		exist[v[0]*60001+v[1]] = true
	}
	res := 0
	for _, v := range commands {
		if v == -1 {
			direIdx += 1
		} else if v == -2 {
			direIdx += 3
		} else {
			for i := 0; i < v; i++ {
				if exist[(px+directions[direIdx%4][0])*60001+py+directions[direIdx%4][1]] {
					break
				}
				px += directions[direIdx%4][0]
				py += directions[direIdx%4][1]
				res = max(res, px*px+py*py)
			}
		}
	}
	return res
}
