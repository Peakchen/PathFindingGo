package finders

import "go-PathFinding/core"

type FinderBase interface {
	FindPath(startX, startY, endX, endY int, grid *core.TGrid) core.DoubleInt32
}
