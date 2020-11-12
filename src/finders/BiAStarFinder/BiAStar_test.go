package BiAStarFinder

import (
	"go-PathFinding/core"
	"go-PathFinding/finders/config"
	"testing"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"
)

func TestBiAStarFinder(t *testing.T) {
	akLog.FmtPrintln("begin BiAStarFinder test...")
	now := time.Now()
	for _, item := range config.PathData {
		itemnow := time.Now()
		grid := core.Grid(len(item.Matrix[0]), len(item.Matrix), item.Matrix)
		opt := &core.Opt{
			AllowDiagonal:    false,
			DontCrossCorners: false,
			DiagonalMovement: core.Never, //Always,Never
			Heuristic:        nil,
			Weight:           0,
		}
		finder := CreateBiAStarFinder(opt)
		result := finder.FindPath(item.StartX, item.StartY, item.EndX, item.EndY, grid)
		akLog.FmtPrintln("result: ", result, "\n", float64(time.Since(itemnow).Nanoseconds())/float64(1e9))
	}
	akLog.FmtPrintln("spend: ", float64(time.Since(now).Nanoseconds())/float64(1e9))
}
