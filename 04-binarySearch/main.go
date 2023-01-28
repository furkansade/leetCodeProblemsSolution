func search(nums []int, target int) int {
	for i, _ := range nums {
		if nums[i] == target {
			return i
		}
	}
	return -1
}