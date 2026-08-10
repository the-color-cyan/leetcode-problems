func twoSum(nums []int, target int) []int {
	nMap := make(map[int]int)

	for i, n := range nums {
		diff := target - n
		diffIndex, ok := nMap[diff]

		if ok {
			return []int{diffIndex, i}
		}

		nMap[n] = i
	}

	return []int{}
}
