func containsDuplicate(nums []int) bool {
	m := make(map[int]int)

	for _, v := range nums {

		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			return true
		}

		// _, ok := m[v]
		// if ok == false {
		// 	m[v] = 1
		// } else {
		// 	return true
		// }

	}
	return false
}