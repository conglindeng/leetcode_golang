package middle

func MaxSubArray(nums []int) int {
	res := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		sum += v
		res = max(res, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return res
} 

func MaxSubArray_devide_conquer(nums []int) int {
	// var deviceAndConquer func(i, j int, nums []int) 
	// deviceAndConquer = func(i, j int, nums []int)  {

	// }
	
	//todo:dcl
	return 0
}
