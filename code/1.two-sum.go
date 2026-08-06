package leetcode

func twoSum(nums []int, target int) []int {
	nMap := make(map[int]int)

	for i, n := range nums {
		diff := target - n
		diffIndex := nMap[diff]

		if diffIndex != nil {
			return [2]int{diffIndex, i}
		}

		nMap[n] = i
	}
}
