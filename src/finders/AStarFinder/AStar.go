package AStarFinder

/*
	by stefan 2572915286@qq.com
	Based upon https://github.com/qiao/PathFinding.js
*/

import (
	"go-PathFinding/core"
	"math"
)

type TAStarFinder struct {
	FinderOpt *core.Opt
}

/**
* A* path-finder.
  Based upon https://github.com/bgrins/javascript-astar
* @constructor
* @param {Object} opt
* @param {boolean} opt.allowDiagonal Whether diagonal movement is allowed.
*     Deprecated, use diagonalMovement instead.
* @param {boolean} opt.dontCrossCorners Disallow diagonal movement touching
*     block corners. Deprecated, use diagonalMovement instead.
* @param {DiagonalMovement} opt.diagonalMovement Allowed diagonal movement.
* @param {function} opt.heuristic Heuristic function to estimate the distance
*     (defaults to manhattan).
* @param {number} opt.weight Weight to apply to the heuristic to allow for
*     suboptimal paths, in order to speed up the search.
*/

func CreateAStarFinder(opt *core.Opt) (this *TAStarFinder) {
	this = &TAStarFinder{
		FinderOpt: opt,
	}
	if opt.Heuristic == nil {
		this.FinderOpt.Heuristic = core.Manhattan
	}
	if opt.Weight == 0 {
		this.FinderOpt.Weight = 1
	}

	if this.FinderOpt.DiagonalMovement == 0 {
		if !this.FinderOpt.AllowDiagonal {
			this.FinderOpt.DiagonalMovement = core.Never
		} else {
			if this.FinderOpt.DontCrossCorners {
				this.FinderOpt.DiagonalMovement = core.OnlyWhenNoObstacles
			} else {
				this.FinderOpt.DiagonalMovement = core.IfAtMostOneObstacle
			}
		}
	}

	// When diagonal movement is allowed the Manhattan heuristic is not
	//admissible. It should be octile instead
	if opt.Heuristic != nil {
		this.FinderOpt.Heuristic = opt.Heuristic
	} else {
		if this.FinderOpt.DiagonalMovement == core.Never {
			this.FinderOpt.Heuristic = core.Manhattan
		} else {
			this.FinderOpt.Heuristic = core.Octile
		}
	}
	return
}

/**
 * Find and return the the path.
 * @return {core.DoubleInt32} The path, including both start and
 *     end positions.
 */
func (this *TAStarFinder) FindPath(startX, startY, endX, endY int, grid *core.TGrid) core.DoubleInt32 {
	var path = core.DoubleInt32{}

	var openList = core.NewGridHeap()
	var startNode = &core.AStarGrid{
		TNode:  grid.GetNodeAt(startX, startY),
		F:      0.0,
		G:      0.0,
		H:      0,
		Opened: false,
		Closed: false,
	}
	endNode := grid.GetNodeAt(endX, endY)
	heuristic := this.FinderOpt.Heuristic
	diagonalMovement := this.FinderOpt.DiagonalMovement
	weight := this.FinderOpt.Weight

	// var node, neighbor *core.AStarGrid
	// var neighbors core.ArrayNode
	//var i, l int
	var x, y int32
	var ng float64

	// set the `g` and `f` value of the start node to be 0
	startNode.G = 0.0
	startNode.F = 0.0

	// push the start node into the open list
	openList.Push(startNode)
	startNode.Opened = true

	// record walked positions
	var walkedMap = map[string]bool{}

	// while the open list is not empty
	for !openList.Empty() {
		// pop the position of node which has the minimum `f` value.
		node := openList.Pop()
		node.Closed = true
		walkedMap[core.NodeGroupStr(node.X, node.Y)] = true

		// if reached the end position, construct the path and return it
		if node.IsEqual(endNode) {
			return core.Backtrace(endNode)
		}

		// get neigbours of the current node
		neighbors := grid.GetNeighbors(node.TNode, diagonalMovement)
		for i := 0; i < len(neighbors); i++ {
			if walkedMap[core.NodeGroupStr(neighbors[i].X, neighbors[i].Y)] {
				continue
			}

			neighbor := &core.AStarGrid{
				TNode:  neighbors[i],
				F:      0.0,
				G:      0.0,
				H:      0,
				Opened: false,
				Closed: false,
			}

			if neighbor.Closed {
				continue
			}

			x = neighbor.X
			y = neighbor.Y

			// get the distance between current node and the neighbor
			// and calculate the next g score
			if x-node.X == 0 || y-node.Y == 0 {
				ng = node.G + float64(1)
			} else {
				ng = node.G + core.SQRT2
			}

			// check if the neighbor has not been inspected yet, or
			// can be reached with smaller cost from the current node
			if !neighbor.Opened || ng < neighbor.G {
				neighbor.G = ng
				if neighbor.H == 0 {
					neighbor.H = weight * heuristic(int32(math.Abs(float64(x-int32(endX)))), int32(math.Abs(float64((y-int32(endY))))))
				}
				neighbor.F = neighbor.G + float64(neighbor.H)
				neighbor.Parent = node.TNode

				if !neighbor.Opened {
					openList.Push(neighbor)
					neighbor.Opened = true
				} else {
					// the neighbor can be reached with smaller cost.
					// Since its f value has been updated, we have to
					// update its position in the open list
					openList.UpdateItem(neighbor)
				}
			}
		} // end for each neighbor
	} // end while not open list empty

	// fail to find the path
	return path
}
