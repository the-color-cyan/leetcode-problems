func missingInteger(nums []int) int {
	if len(nums) == 1 {
		return nums[0] + 1
	}

	end, sum := findSeqPrefix(nums)
	remaining := mapRemaining(nums, end)

	v := sum
	for remaining.Contains(v) {
		v++
	}

	return v
}

// end is exclusive
func findSeqPrefix(nums []int) (end int, sum int) {
	if len(nums) < 2 {
		panic(len(nums))
	}

	sum = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			break
		}

		sum += nums[i]
		end = i + 1
	}

	return
}

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(val T) {
	s[val] = struct{}{}
}

func (s Set[T]) Contains(val T) bool {
	_, exists := s[val]
	return exists
}

func mapRemaining(nums []int, end int) Set[int] {
	remaining := make(Set[int])

	for _, n := range nums[end:] {
		remaining.Add(n)
	}

	return remaining
}
