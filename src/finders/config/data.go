package config

import (
	"go-PathFinding/core"
)

var (
	PathData = []struct {
		StartX         int
		StartY         int
		EndX           int
		EndY           int
		Matrix         core.DoubleInt32
		ExpectedLength int
	}{}
)

func init() {
	PathData = append(PathData, struct {
		StartX         int
		StartY         int
		EndX           int
		EndY           int
		Matrix         core.DoubleInt32
		ExpectedLength int
	}{
		StartX: 0,
		StartY: 0,
		EndX:   1,
		EndY:   1,
		Matrix: core.DoubleInt32{
			{0, 0},
			{1, 0}},
		ExpectedLength: 3,
	})

	PathData = append(PathData, struct {
		StartX         int
		StartY         int
		EndX           int
		EndY           int
		Matrix         core.DoubleInt32
		ExpectedLength int
	}{
		StartX: 0,
		StartY: 0,
		EndX:   2,
		EndY:   2,
		Matrix: core.DoubleInt32{
			{0, 0, 0},
			{1, 1, 0},
			{0, 0, 0}},
		ExpectedLength: 3,
	})

	PathData = append(PathData, struct {
		StartX         int
		StartY         int
		EndX           int
		EndY           int
		Matrix         core.DoubleInt32
		ExpectedLength int
	}{
		StartX: 1,
		StartY: 1,
		EndX:   4,
		EndY:   4,
		Matrix: core.DoubleInt32{
			{0, 0, 0, 0, 0},
			{1, 0, 1, 1, 0},
			{1, 0, 1, 0, 0},
			{0, 1, 0, 0, 0},
			{1, 0, 1, 1, 0},
			{0, 0, 1, 0, 0},
		},
		ExpectedLength: 9,
	})

	PathData = append(PathData, struct {
		StartX         int
		StartY         int
		EndX           int
		EndY           int
		Matrix         core.DoubleInt32
		ExpectedLength int
	}{
		StartX: 0,
		StartY: 3,
		EndX:   3,
		EndY:   3,
		Matrix: core.DoubleInt32{
			{0, 0, 0, 0, 0},
			{0, 0, 1, 1, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 1, 0, 0},
			{1, 0, 1, 1, 0},
			{0, 0, 0, 0, 0},
		},
		ExpectedLength: 10,
	})
}
