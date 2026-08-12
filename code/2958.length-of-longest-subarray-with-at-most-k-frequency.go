func maxSubarrayLength(nums []int, k int) int {
	freq := map[int]int{}
	longest, left := 0, 0

	for right := range nums {
		freq[nums[right]] += 1

		for freq[nums[right]] > k {
			freq[nums[left]] -= 1
			left++
		}

		if length := right + 1 - left; length > longest {
			longest = length
		}
	}

	return longest
}
