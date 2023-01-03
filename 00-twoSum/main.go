func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, v := range nums {
		val, found := hashMap[target-v]
		if found {
			return []int{val, i}
		}
		// if found == true {}

		hashMap[v] = i
	}

	return nil
}
