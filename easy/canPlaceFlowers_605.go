package easy

func canPlaceFlowers(flowerbed []int, n int) bool {
	pre := -1
	i := 0
	count := 0
	for i < len(flowerbed) {
		if flowerbed[i] == 1 {
			count += (i - pre - 1) / 2
			pre = i + 1
		}
		i++
	}
	if pre < len(flowerbed)-1 {
		count += (i - pre) / 2
	}
	return count >= n
}
