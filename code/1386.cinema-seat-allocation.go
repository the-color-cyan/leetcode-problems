package leetcode

const (
	blockSize   int = 4
	blocks      int = 3
	seatsPerRow int = 10
)

var seatBlocks = [blocks]SeatBlock{
	{2, 3, 4, 5},
	{4, 5, 6, 7},
	{6, 7, 8, 9},
}

type SeatBlock [blockSize]int

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

func groupFitMap(seatBlocks []SeatBlock) map[int]int {
	fitMap := make(map[int]int)

	for _, block := range seatBlocks {
		for _, seat := range block {
		}
	}
}

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
}
