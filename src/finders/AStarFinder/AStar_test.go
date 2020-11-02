package AStarFinder

import (
	"fmt"
	"go-PathFinding/core"
	"go-PathFinding/finders/config"
	"testing"
)

func TestAStarFinder(t *testing.T) {
	for _, item := range config.PathData {
		grid := core.Grid(len(item.Matrix[0]), len(item.Matrix), item.Matrix)
		opt := &core.Opt{
			AllowDiagonal:    false,
			DontCrossCorners: false,
			DiagonalMovement: core.Always,
			Heuristic:        nil,
			Weight:           0,
		}
		finder := CreateAStarFinder(opt)
		result := finder.FindPath(item.StartX, item.StartY, item.EndX, item.EndY, grid)
		fmt.Println("result: ", result)
	}
}
