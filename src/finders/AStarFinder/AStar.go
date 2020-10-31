package AStarFinder

import (
	"go-PathFinding/core"
	"math"
)

type TAStarFinder struct {
	FinderOpt *core.Opt
}

/**
 * A* path-finder. Based upon https://github.com/bgrins/javascript-astar
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

func CreateAStarFinder(opt *core.Opt) {
	this := &TAStarFinder{
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
}

/**
 * Find and return the the path.
 * @return {core.DoubleInt32} The path, including both start and
 *     end positions.
 */
func (this *TAStarFinder) FindPath(startX, startY, endX, endY int, grid *core.TGrid) core.DoubleInt32 {
	var path = core.DoubleInt32{}

	var openList = NewGridHeap()
	var startNode = &AStarGrid{
		TNode:  grid.GetNodeAt(startX, startY),
		f:      0.0,
		g:      0.0,
		h:      0,
		opened: false,
		closed: false,
	}
	endNode := grid.GetNodeAt(endX, endY)
	heuristic := this.FinderOpt.Heuristic
	diagonalMovement := this.FinderOpt.DiagonalMovement
	weight := this.FinderOpt.Weight

	var node, neighbor *AStarGrid
	var neighbors core.ArrayNode
	//var i, l int
	var x, y int32
	var ng float64

	// set the `g` and `f` value of the start node to be 0
	startNode.g = 0.0
	startNode.f = 0.0

	// push the start node into the open list
	openList.Push(startNode)
	startNode.opened = true

	// while the open list is not empty
	for !openList.Empty() {
		// pop the position of node which has the minimum `f` value.
		node = openList.Pop()
		node.closed = true

		// if reached the end position, construct the path and return it
		if node.IsEqual(endNode) {
			return core.Backtrace(endNode)
		}

		// get neigbours of the current node
		neighbors = grid.GetNeighbors(node.TNode, diagonalMovement)
		for i := 0; i < len(neighbors); i++ {
			neighbor = &AStarGrid{
				TNode:  neighbors[i],
				f:      0.0,
				g:      0.0,
				h:      0,
				opened: false,
				closed: false,
			}

			if neighbor.closed {
				continue
			}

			x = neighbor.X
			y = neighbor.Y

			// get the distance between current node and the neighbor
			// and calculate the next g score
			if x-node.X == 0 || y-node.Y == 0 {
				ng = node.g + float64(1)
			} else {
				ng = node.g + core.SQRT2
			}

			// check if the neighbor has not been inspected yet, or
			// can be reached with smaller cost from the current node
			if !neighbor.opened || ng < neighbor.g {
				neighbor.g = ng
				if neighbor.h == 0 {
					neighbor.h = weight * heuristic(int32(math.Abs(float64(x-int32(endX)))), int32(math.Abs(float64((y-int32(endY))))))
				}
				neighbor.f = neighbor.g + float64(neighbor.h)
				neighbor.Parent = node.TNode

				if !neighbor.opened {
					openList.Push(neighbor)
					neighbor.opened = true
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
