package easy

func twoSum(nums []int, target int) []int {
	num2Index := map[int]int{}
	for i, n := range nums {
		if _, exsit := num2Index[target-n]; exsit {
			return []int{num2Index[target-n], i}
		}
		num2Index[n] = i
	}
	return nil
}
