type Interval struct {
	start int
	end   int
}

type Prefix []int

type Memo [][]int

func (i Interval) Len() int {
	return i.end - i.start
}

func (i Interval) Sum(p Prefix) int {
	return p[i.end] - p[i.start]
}

func NewPrefix(arr []int) Prefix {
	prefix := make([]int, len(arr)+1)

	for i, v := range arr {
		prefix[i+1] = prefix[i] + v
	}

	return prefix
}

func NewMemo(l int) Memo {
	size := l + 1
	memo := make([][]int, size)

	for i := range memo {
		memo[i] = make([]int, size)

		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return memo
}

func (m Memo) Set(i Interval, v int) {
	m[i.start][i.end] = v
}

func (m Memo) Get(i Interval) (int, bool) {
	v := m[i.start][i.end]

	if v == -1 {
		return v, false
	}

	return v, true
}

func stoneGameV(stoneValue []int) int {
	full := Interval{0, len(stoneValue)}
	memo := NewMemo(len(stoneValue))
	prefix := NewPrefix(stoneValue)

	var best func(Interval) int

	scoreAtCut := func(i Interval, cut int) int {
		left, right := Interval{i.start, cut}, Interval{cut, i.end}
		leftSum, rightSum := left.Sum(prefix), right.Sum(prefix)

		switch {
		case leftSum > rightSum:
			return rightSum + best(right)

		case rightSum > leftSum:
			return leftSum + best(left)

		default:
			leftBest := best(left)
			rightBest := best(right)

			return leftSum + max(leftBest, rightBest)
		}
	}

	best = func(i Interval) int {
		length := i.Len()

		if length < 1 {
			panic("invalid interval")
		}

		memoScore, ok := memo.Get(i)
		if ok {
			return memoScore
		}

		if length == 1 {
			memo.Set(i, 0)
			return 0
		}

		score := 0

		for cut := i.start + 1; cut < i.end; cut++ {
			candidate := scoreAtCut(i, cut)
			score = max(candidate, score)
		}

		memo.Set(i, score)
		return score
	}

	return best(full)
}
