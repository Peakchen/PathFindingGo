package AStarFinder

import (
	"go-PathFinding/core"
	"go-PathFinding/finders/config"
	"testing"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"
)

func TestAStarFinder(t *testing.T) {
	now := time.Now()
	for _, item := range config.PathData {
		grid := core.Grid(len(item.Matrix[0]), len(item.Matrix), item.Matrix)
		opt := &core.Opt{
			AllowDiagonal:    false,
			DontCrossCorners: false,
			DiagonalMovement: core.Never, //Aways
			Heuristic:        nil,
			Weight:           0,
		}
		finder := CreateAStarFinder(opt)
		result := finder.FindPath(item.StartX, item.StartY, item.EndX, item.EndY, grid)
		akLog.FmtPrintln("result: ", result)
	}
	akLog.FmtPrintln("spend: ", float64(time.Since(now).Nanoseconds())/float64(1e9))
}
