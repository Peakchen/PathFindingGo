package core

type DoubleInt32 [][]int32
type DoubleInt64 [][]int64

type ArrayInt8 []int8
type ArrayInt16 []int16
type ArrayInt32 []int32
type ArrayInt64 []int64

type ArrayUInt8 []uint8
type ArrayUInt16 []uint16
type ArrayUInt32 []uint32
type ArrayUInt64 []uint64

const (
	SQRT2 = 1.4142135623730951
)

type Opt struct {
	AllowDiagonal    bool
	DontCrossCorners bool
	DiagonalMovement DiagonalMovement
	Heuristic        func(x, y int32) int32
	Weight           int32
}
