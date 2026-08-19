package leetcode

const (
	BlockSize int = 4
	Blocks    int = 3
)

var SeatBlocks = [Blocks]SeatBlock{
	{2, 3, 4, 5},
	{4, 5, 6, 7},
	{6, 7, 8, 9},
}

type SeatBlock [BlockSize]int

// type Set[T comparable] map[T]struct{}
//
// func (s Set[T]) Add(val T) {
// 	s[val] = struct{}{}
// }
//
// func (s Set[T]) Contains(val T) bool {
// 	_, exists := s[val]
// 	return exists
// }

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
}
