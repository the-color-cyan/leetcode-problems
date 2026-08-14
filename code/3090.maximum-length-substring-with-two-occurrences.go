package leetcode

const MaxOccurences int = 2

func maximumLengthSubstring(s string) int {
	// chars[b] = chars (1-based) index byte b was seen at in s
	// alternatively: the index after the chars occurence of b in s
	var chars Characters
	longest, start := 0, 0

	for end := range s {
		char := chars[s[end]]

		if char.last[0] > start {
			start = char.last[0]
		}

		char.last[] = end + 1

		if length := end + 1 - start; length > longest {
			longest = length
		}
	}

	return longest
}

type CharInfo struct {
	last [MaxOccurences]int // last x indices the character occured at
}

type Characters [256]CharInfo
