package BiAStarFinder

/*
	by stefan 2572915286@qq.com
	Based upon https://github.com/qiao/PathFinding.js
*/

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
		TNode:  grid.GetNodeAt(endX, endY),
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
	var walkedMap = map[string]*core.AStarGrid{}

	//check produce repeat parent.
	checkNodeParentSame := func(src, dst *core.TNode) bool {
		srcParent := src.Parent
		for srcParent != nil {
			if core.NodeGroupStr(src.Parent.X, src.Parent.Y) == core.NodeGroupStr(dst.X, dst.Y) {
				return true
			}
			srcParent = srcParent.Parent
		}
		return false
	}

	//check and sort path right.
	sortPathOrder := func(src, dst *core.TNode) core.DoubleInt32 {
		result := core.BiBacktrace(src, dst)
		fistCoord := core.Array2Coordinate(result[0])
		lastCoord := core.Array2Coordinate(result[len(result)-1])
		if (fistCoord.X != startNode.TNode.X && fistCoord.Y != startNode.TNode.Y) &&
			(lastCoord.X != endNode.TNode.X && lastCoord.Y != endNode.TNode.Y) {
			core.Reverse(result)
		} else if !(fistCoord.X == startNode.TNode.X && fistCoord.Y == startNode.TNode.Y &&
			lastCoord.X == endNode.TNode.X && lastCoord.Y == endNode.TNode.Y) {
			if fistCoord.X == lastCoord.X && fistCoord.Y == lastCoord.Y &&
				(startOpenList.Empty() || endOpenList.Empty()) {
				result = result[:len(result)-1]
			} else {
				result = nil
			}
		}
		return result
	}

	//check parent find end node.
	isLinkEndNode := func(node *core.TNode) bool {
		if node.Parent == nil {
			return false
		}
		return node.Parent.IsEqual(startNode.TNode) && node.IsEqual(endNode.TNode)
	}

	getPathNodes := func(list *core.GridHeap, endflag int, openflag int) core.DoubleInt32 {
		// pop the position of start node which has the minimum `f` value.
		node := list.Pop()
		node.Closed = false

		// for short distance path find.
		if isLinkEndNode(node.TNode) {
			return core.DoubleInt32{
				core.ArrayInt32{startNode.X, startNode.Y},
				core.ArrayInt32{endNode.X, endNode.Y}}
		}

		walkedMap[core.NodeGroupStr(node.X, node.Y)] = node

		// get neigbours of the current node
		neighbors := grid.GetNeighbors(node.TNode, diagonalMovement)
		for i := 0; i < len(neighbors); i++ {
			if checkNodeParentSame(node.TNode, neighbors[i]) {
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

			walkedNode := walkedMap[core.NodeGroupStr(neighbors[i].X, neighbors[i].Y)]
			if walkedNode != nil {
				neighbor = walkedNode
			}

			if neighbor.Openedflag == endflag {
				return sortPathOrder(node.TNode, neighbor.TNode)
			}

			if neighbor.Closed {
				continue
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

				if neighbor.Parent == nil {
					neighbor.Parent = node.TNode
				}

				if neighbor.Openedflag == 0 {
					list.Push(neighbor)
					neighbor.Openedflag = openflag
				} else {
					// the neighbor can be reached with smaller cost.
					// Since its f value has been updated, we have to
					// update its position in the open list
					list.UpdateItem(neighbor)
				}
			}
		} // end for each neighbor
		// fail to find the path
		return core.DoubleInt32{}
	}

	// while both the open lists are not empty
	for !startOpenList.Empty() && !endOpenList.Empty() {

		startRet := getPathNodes(startOpenList, BY_END, BY_START)
		if len(startRet) != 0 {
			return startRet
		}

		endRet := getPathNodes(endOpenList, BY_START, BY_END)
		if len(endRet) != 0 {
			return endRet
		}
	} // end while not open list empty

	// fail to find the path
	return core.DoubleInt32{}
}
