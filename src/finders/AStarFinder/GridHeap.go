package AStarFinder

import (
	"go-PathFinding/core"
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
	*core.TNode
	f      float64
	g      float64
	h      int32
	opened bool
	closed bool
}

func NewGridHeap() *GridHeap {
	return &GridHeap{
		grids: []*AStarGrid{},
	}
}

func (this *GridHeap) autoSort() {
	sort.Slice(this.grids, func(i, j int) bool {
		return this.grids[i].f < this.grids[j].f
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
