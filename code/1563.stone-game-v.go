func stoneGameV(stoneValue []int) int {
	i := Interval{0, len(stoneValue)}

	return bestScore(i, stoneValue)
}

type Interval struct {
	start int
	end   int
}

func (i Interval) Len() int {
	return i.end - i.start
}

type Memo [][]int

func NewMemo(l int) Memo {
	memo := make([][]int, l)
	for i := range memo {
		memo[i] = make([]int, l)

		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return memo
}

func (m Memo) Add(i Interval, v int) {
	m[i.start][i.end] = v
}

func (m Memo) Get(i Interval) (int, bool) {
	v := m[i.start][i.end]

	if v == -1 {
		return v, false
	}

	return v, true
}

func bestScore(i Interval, arr []int) int {
	memo := NewMemo(len(arr) + 1)
	prefix := prefixSums(arr)

	var best func(Interval) int

	best = func(i Interval) int {
		length := i.Len()

		if length < 1 {
			panic("invalid interval")
		}

		score, ok := memo.Get(i)
		if ok {
			return score
		}

		if length == 1 {
			memo.Add(i, 0)
			return 0
		}

		score = 0

		for cut := i.start + 1; cut < i.end; cut++ {
			left, right := Interval{i.start, cut}, Interval{cut, i.end}
			leftSum, rightSum := intervalSum(left, prefix), intervalSum(right, prefix)

			var candidate int

			switch {
			case leftSum > rightSum:
				candidate = rightSum + best(right)

			case rightSum > leftSum:
				candidate = leftSum + best(left)

			default:
				leftBest := best(left)
				rightBest := best(right)

				candidate = leftSum + max(leftBest, rightBest)
			}

			score = max(candidate, score)
		}

		memo.Add(i, score)
		return score
	}

	return best(i)
}

func intervalSum(i Interval, prefix []int) (sum int) {
	return prefix[i.end] - prefix[i.start]
}

func prefixSums(arr []int) []int {
	prefix := make([]int, len(arr)+1)

	for i, v := range arr {
		prefix[i+1] = prefix[i] + v
	}

	return prefix
}
