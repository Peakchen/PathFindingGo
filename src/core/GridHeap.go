package core

import (
	"sort"
)

/*
	by stefan 2572915286@qq.com
	Based upon https://github.com/qiao/PathFinding.js
*/

type GridHeap struct {
	grids []*AStarGrid
}

type AStarGrid struct {
	*TNode
	F          float64
	G          float64
	H          int32
	Opened     bool
	Closed     bool
	Openedflag int // another opened used
}

func NewGridHeap() *GridHeap {
	return &GridHeap{
		grids: []*AStarGrid{},
	}
}

func (this *GridHeap) autoSort() {
	sort.Slice(this.grids, func(i, j int) bool {
		return this.grids[i].F < this.grids[j].F
	})
}

func (this *GridHeap) Push(new *AStarGrid) {
	this.grids = append(this.grids, new)
	//this.autoSort()
}

func (this *GridHeap) Pop() (grid *AStarGrid) {
	total := len(this.grids)
	if total == 0 {
		return nil
	}
	grid = this.grids[total-1]
	this.grids = this.grids[:total-1]
	return
}

func (this *GridHeap) Empty() bool {
	return len(this.grids) == 0
}

func (this *GridHeap) UpdateItem(grid *AStarGrid) {
	i := sort.Search(len(this.grids), func(i int) bool { return this.grids[i].X == grid.X && this.grids[i].Y == grid.X })
	if i < len(this.grids) && (this.grids[i].X == grid.X && this.grids[i].Y == grid.Y) {
		this.grids[i] = grid
	}
}
