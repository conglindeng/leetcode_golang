package easy

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	left := min(start, destination)
	right := max(start, destination)
	sum := 0
	for i := 0; i < len(distance); i++ {
		sum += distance[i]
	}
	cost := 0
	for i := left; i < right; i++ {
		cost += distance[i]
	}
	return min(cost, sum-cost)
}
