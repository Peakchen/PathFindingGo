package core

import (
	"github.com/Peakchen/xgameCommon/akLog"
)

type TGrid struct {
	width  int
	height int
	nodes  DoubleNode
}

/**
 * The Grid class, which serves as the encapsulation of the layout of the nodes.
 * @constructor
 * @param {number} width Number of columns of the grid, or matrix
 * @param {number} height Number of rows of the grid.
 * @param {Doubleint32} [matrix] - A 0-1 matrix
 *     representing the walkable status of the nodes(0 or false for walkable).
 *     If the matrix is not supplied, all the nodes will be walkable.  */

func Grid(width, height int, matrix DoubleInt32) *TGrid {
	return &TGrid{
		width:  width,
		height: height,
		nodes:  buildNodes(width, height, matrix),
	}
}

/**
 * Build and return the nodes.
 * @private
 * @param {number} width
 * @param {number} height
 * @param {DoubleNode} [matrix] - A 0-1 matrix representing
 *     the walkable status of the nodes.
 * @see Grid
 */

func buildNodes(width, height int, matrix DoubleInt32) DoubleNode {
	var nodes = make(DoubleNode, height)

	for i := 0; i < height; i++ {
		nodes[i] = make(ArrayNode, width)
		for j := 0; j < width; j++ {
			nodes[i][j] = Node(int32(j), int32(i), true)
		}
	}

	if matrix == nil {
		return nodes
	}

	if len(matrix) != height || len(matrix[0]) != width {
		akLog.Error("Matrix size does not fit")
		return nodes
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if nodes[i][j] == nil {
				panic("invalid node.")
			}
			// 0, false, null will be walkable
			// while others will be un-walkable
			nodes[i][j].walkable = false
		}
	}

	return nodes
}

func (this *TGrid) getNodeAt(x, y int) *TNode {
	return this.nodes[y][x]
}

/**
 * Determine whether the node at the given position is walkable.
 * (Also returns false if the position is outside the grid.)
 * @param {number} x - The x coordinate of the node.
 * @param {number} y - The y coordinate of the node.
 * @return {boolean} - The walkability of the node.
 */
func (this *TGrid) IsWalkableAt(x, y int) bool {
	return this.isInside(x, y) && this.nodes[y][x].walkable
}

/**
 * Determine whether the position is inside the grid.
 * XXX: `grid.isInside(x, y)` is wierd to read.
 * It should be `(x, y) is inside grid`, but I failed to find a better
 * name for this method.
 * @param {number} x
 * @param {number} y
 * @return {boolean}
 */
func (this *TGrid) isInside(x, y int) bool {
	return (x >= 0 && x < this.width) && (y >= 0 && y < this.height)
}

/**
 * Set whether the node on the given position is walkable.
 * NOTE: throws exception if the coordinate is not inside the grid.
 * @param {number} x - The x coordinate of the node.
 * @param {number} y - The y coordinate of the node.
 * @param {boolean} walkable - Whether the position is walkable.
 */
func (this *TGrid) setWalkableAt(x, y int32, walkable bool) {
	this.nodes[y][x].walkable = walkable
}

/**
 * Get the neighbors of the given node.
 *
 *     offsets      diagonalOffsets:
 *  +---+---+---+    +---+---+---+
 *  |   | 0 |   |    | 0 |   | 1 |
 *  +---+---+---+    +---+---+---+
 *  | 3 |   | 1 |    |   |   |   |
 *  +---+---+---+    +---+---+---+
 *  |   | 2 |   |    | 3 |   | 2 |
 *  +---+---+---+    +---+---+---+
 *
 *  When allowDiagonal is true, if offsets[i] is valid, then
 *  diagonalOffsets[i] and
 *  diagonalOffsets[(i + 1) % 4] is valid.
 * @param {Node} node
 * @param {DiagonalMovement} diagonalMovement
 */
func (this *TGrid) getNeighbors(node *TNode, move DiagonalMovement) ArrayNode {
	var x = int(node.x)
	var y = int(node.y)
	var neighbors = ArrayNode{}
	var (
		s0    = false
		d0    = false
		s1    = false
		d1    = false
		s2    = false
		d2    = false
		s3    = false
		d3    = false
		nodes = this.nodes
	)

	// ↑
	if this.IsWalkableAt(x, y-1) {
		neighbors = append(neighbors, nodes[y-1][x])
		s0 = true
	}
	// →
	if this.IsWalkableAt(x+1, y) {
		neighbors = append(neighbors, nodes[y][x+1])
		s1 = true
	}
	// ↓
	if this.IsWalkableAt(x, y+1) {
		neighbors = append(neighbors, nodes[y+1][x])
		s2 = true
	}
	// ←
	if this.IsWalkableAt(x-1, y) {
		neighbors = append(neighbors, nodes[y][x-1])
		s3 = true
	}

	if move == Never {
		return neighbors
	}

	if move == OnlyWhenNoObstacles {
		d0 = s3 && s0
		d1 = s0 && s1
		d2 = s1 && s2
		d3 = s2 && s3
	} else if move == IfAtMostOneObstacle {
		d0 = s3 || s0
		d1 = s0 || s1
		d2 = s1 || s2
		d3 = s2 || s3
	} else if move == Always {
		d0 = true
		d1 = true
		d2 = true
		d3 = true
	} else {
		panic("Incorrect value of diagonalMovement")
	}

	// ↖
	if d0 && this.IsWalkableAt(x-1, y-1) {
		neighbors = append(neighbors, nodes[y-1][x-1])
	}
	// ↗
	if d1 && this.IsWalkableAt(x+1, y-1) {
		neighbors = append(neighbors, nodes[y-1][x+1])
	}
	// ↘
	if d2 && this.IsWalkableAt(x+1, y+1) {
		neighbors = append(neighbors, nodes[y+1][x+1])
	}
	// ↙
	if d3 && this.IsWalkableAt(x-1, y+1) {
		neighbors = append(neighbors, nodes[y+1][x-1])
	}

	return neighbors
}

/**
 * Get a clone of this grid.
 * @return {Grid} Cloned grid.
 */
func (this *TGrid) clone() *TGrid {
	var i, j int
	width := this.width
	height := this.height
	thisNodes := this.nodes

	newGrid := Grid(width, height, nil)
	newNodes := make(DoubleNode, height)

	for i = 0; i < height; i++ {
		newNodes[i] = make(ArrayNode, width)
		for j = 0; j < width; j++ {
			newNodes[i][j] = Node(int32(j), int32(i), thisNodes[i][j].walkable)
		}
	}

	newGrid.nodes = newNodes

	return newGrid
}
