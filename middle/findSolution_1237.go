package middle

import (
	"sort"
)

/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

func FindSolution(customFunction func(int, int) int, z int) [][]int {
	res := make([][]int, 0)
	for x := 1; x <= 1000 && x <= z; x++ {
		y := sort.Search(1000, func(m int) bool { return customFunction(x, m) >= z })
		if y > 0 && customFunction(x, y) == z {
			res = append(res, []int{x, y})
		}
	}
	return res
}

func FindSolution_Pointer(customFunction func(int, int) int, z int) [][]int {
	res := make([][]int, 0)
	for x, y := 1, 1000; x <= 1000 && y >= 0; x++ {
		for y > 0 && customFunction(x, y) > z {
			y--
		}
		if y > 0 && customFunction(x, y) == z {
			res = append(res, []int{x, y})
		}
	}
	return res
}
