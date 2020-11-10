package BiAStarFinder

import (
	"go-PathFinding/core"
	"go-PathFinding/finders/AStarFinder"
	"math"
)

type BiAStarFinder struct {
	*AStarFinder.TAStarFinder
}

/**
 * A* path-finder.
 * based upon https://github.com/bgrins/javascript-astar
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

func CreateBiAStarFinder(opt *core.Opt) (this *BiAStarFinder) {
	return &BiAStarFinder{
		TAStarFinder: AStarFinder.CreateAStarFinder(opt),
	}
}

/**
 * Find and return the the path.
 * @return {core.DoubleInt32} The path, including both start and
 *     end positions.
 */
func (this *BiAStarFinder) FindPath(startX, startY, endX, endY int, grid *core.TGrid) core.DoubleInt32 {
	var startOpenList = core.NewGridHeap()
	var endOpenList = core.NewGridHeap()

	var startNode = &core.AStarGrid{
		TNode:  grid.GetNodeAt(startX, startY),
		F:      0.0,
		G:      0.0,
		H:      0,
		Opened: false,
		Closed: false,
	}

	var endNode = &core.AStarGrid{
		TNode:  grid.GetNodeAt(startX, startY),
		F:      0.0,
		G:      0.0,
		H:      0,
		Opened: false,
		Closed: false,
	}

	heuristic := this.FinderOpt.Heuristic
	diagonalMovement := this.FinderOpt.DiagonalMovement
	weight := this.FinderOpt.Weight

	var ng float64
	var BY_START = 1
	var BY_END = 2

	// set the `g` and `f` value of the start node to be 0
	// and push it into the start open list
	startNode.G = 0.0
	startNode.F = 0.0
	startOpenList.Push(startNode)
	startNode.Openedflag = BY_START

	// set the `g` and `f` value of the end node to be 0
	// and push it into the open open list
	endNode.G = 0.0
	endNode.F = 0.0
	endOpenList.Push(endNode)
	endNode.Openedflag = BY_END

	// record walked positions
	var walkedMap = map[string]bool{}

	// while both the open lists are not empty
	for !startOpenList.Empty() && !endOpenList.Empty() {

		// pop the position of start node which has the minimum `f` value.
		node := startOpenList.Pop()
		node.Closed = false

		walkedMap[core.NodeGroupStr(node.X, node.Y)] = true

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

			if neighbor.Openedflag == BY_END {
				return core.BiBacktrace(node.TNode, neighbor.TNode)
			}

			x := neighbor.X
			y := neighbor.Y

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
					startOpenList.Push(neighbor)
					neighbor.Openedflag = BY_START
				} else {
					// the neighbor can be reached with smaller cost.
					// Since its f value has been updated, we have to
					// update its position in the open list
					startOpenList.UpdateItem(neighbor)
				}
			}
		} // end for each neighbor

		// pop the position of end node which has the minimum `f` value.
		node = endOpenList.Pop()
		node.Closed = true

		// get neigbours of the current node
		neighbors = grid.GetNeighbors(node.TNode, diagonalMovement)
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

			if neighbor.Openedflag == BY_START {
				return core.BiBacktrace(node.TNode, neighbor.TNode)
			}

			x := neighbor.X
			y := neighbor.Y

			// get the distance between current node and the neighbor
			// and calculate the next g score
			if x-node.X == 0 || y-node.Y == 0 {
				ng = node.G + float64(1)
			} else {
				ng = node.G + core.SQRT2
			}

			// check if the neighbor has not been inspected yet, or
			// can be reached with smaller cost from the current node
			if neighbor.Openedflag == 0 || ng < neighbor.G {
				neighbor.G = ng
				if neighbor.H == 0 {
					neighbor.H = weight * heuristic(int32(math.Abs(float64(x-int32(endX)))), int32(math.Abs(float64((y-int32(endY))))))
				}
				neighbor.F = neighbor.G + float64(neighbor.H)
				neighbor.Parent = node.TNode

				if neighbor.Openedflag == 0 {
					endOpenList.Push(neighbor)
					neighbor.Openedflag = BY_END
				} else {
					// the neighbor can be reached with smaller cost.
					// Since its f value has been updated, we have to
					// update its position in the open list
					endOpenList.UpdateItem(neighbor)
				}
			}
		} // end for each neighbor
	} // end while not open list empty

	// fail to find the path
	return core.DoubleInt32{}
}
