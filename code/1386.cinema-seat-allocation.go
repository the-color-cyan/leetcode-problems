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

func excludedBlocksInRowBySeatAllocation(seatBlocks []SeatBlock) map[int]int {
	exclusionMap := make(map[int]int)
	goodBlocks := 

	for seat := 1; seat <= seatsPerRow; seat++ {
		for _, block := range seatBlocks {
			if isSeatInBlock(seat, block) {
				goodBlocks--
			}
		}
	}
}

func isSeatInBlock(seat int, block SeatBlock) bool {
	for _, blockSeat := range block {
		if seat == blockSeat {
			return true
		}
	}

	return false
}

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
}
