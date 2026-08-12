func lengthOfLongestSubstring(s string) int {
	// last[b] = last (1-based) index byte b was seen at in s
	// alternatively: the index after the last occurence of b in s
	var last [256]int
	longest, start := 0, 0

	for end := range s {
		b := s[end]

		if last[b] > start {
			start = last[b]
		}

		last[b] = end + 1

		if length := end + 1 - start; length > longest {
			longest = length
		}
	}

	return longest
}
