package core

/*
	by stefan 2572915286@qq.com
	Based upon https://github.com/qiao/PathFinding.js
*/

import "math"

/**
 * Backtrace according to the Parent records and return the path.
 * (including both start and end nodes)
 * @param {Node} node End node
 * @return {DoubleInt32} the path
 */
func Backtrace(node *TNode) DoubleInt32 {
	var path = DoubleInt32{ArrayInt32{node.X, node.Y}}
	for node.Parent != nil {
		node = node.Parent
		path = append(path, ArrayInt32{node.X, node.Y})
	}
	Reverse(path)
	return path
}

/**
 * Backtrace from start and end node, and return the path.
 * (including both start and end nodes)
 * @param {Node}
 * @param {Node}
 */
func biBacktrace(nodeA, nodeB *TNode) DoubleInt32 {
	pathA := Backtrace(nodeA)
	pathB := Backtrace(nodeB)
	Reverse(pathB)
	pathA = append(pathA, pathB...)
	return pathA
}

/**
 * Compute the length of the path.
 * @param {Array<Array<number>>} path The path
 * @return {number} The length of the path
 */
func pathLength(path DoubleInt32) int {
	var i, sum int
	var a, b ArrayInt32
	var dx, dy int32
	for i = 1; i < len(path); i++ {
		a = path[i-1]
		b = path[i]
		dx = a[0] - b[0]
		dy = a[1] - b[1]
		sum += int(math.Ceil(math.Sqrt(float64(dx*dx + dy*dy))))
	}
	return sum
}

/**
 * Given the start and end coordinates, return all the coordinates lying
 * on the line formed by these coordinates, based on Bresenham's algorithm.
 * http://en.wikipedia.org/wiki/Bresenham's_line_algorithm#Simplification
 * @param {number} x0 Start X coordinate
 * @param {number} y0 Start Y coordinate
 * @param {number} x1 End X coordinate
 * @param {number} y1 End Y coordinate
 * @return {DoubleInt32} The coordinates on the line
 */
func interpolate(x0, y0, x1, y1 int32) DoubleInt32 {
	var (
		line            = DoubleInt32{}
		sx, sy          int32
		dx, dy, err, e2 float64
	)

	dx = math.Abs(float64(x1 - x0))
	dy = math.Abs(float64(y1 - y0))

	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}

	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}

	err = dx - dy

	for true {
		line = append(line, []int32{x0, y0})

		if x0 == x1 && y0 == y1 {
			break
		}

		e2 = float64(2) * err
		if e2 > -dy {
			err = err - dy
			x0 = x0 + sx
		}
		if e2 < dx {
			err = err + dx
			y0 = y0 + sy
		}
	}

	return line
}

/**
 * Given a compressed path, return a new path that has all the segments
 * in it interpolated.
 * @param [][]int32 path The path
 * @return [][]int32 expanded path
 */

func expandPath(path DoubleInt32) DoubleInt32 {
	var (
		expanded       = DoubleInt32{}
		pathlen        = len(path)
		coord0, coord1 ArrayInt32
	)

	if pathlen < 2 {
		return expanded
	}

	for i := 0; i < pathlen-1; i++ {
		coord0 = path[i]
		coord1 = path[i+1]

		interpolated := interpolate(coord0[0], coord0[1], coord1[0], coord1[1])
		interpolatedLen := len(interpolated)
		for j := 0; j < interpolatedLen-1; j++ {
			expanded = append(expanded, interpolated[j])
		}
	}
	expanded = append(expanded, path[pathlen-1])
	return expanded
}

/**
 * Smoothen the give path.
 * The original path will not be modified; a new path will be returned.
 * @param {PF.Grid} grid
 * @param {DoubleInt32} path The path
 */

func smoothenPath(grid *TGrid, path DoubleInt32) DoubleInt32 {
	var pathlen = len(path)
	var x0 = path[0][0]         // path start X
	var y0 = path[0][1]         // path start Y
	var x1 = path[pathlen-1][0] // path end X
	var y1 = path[pathlen-1][1] // path end Y
	var sx, sy int32            // current start coordinate
	var ex, ey int32            // current end coordinate
	var newPath, line DoubleInt32
	var i, j int
	var blocked bool
	var testCoord, coord ArrayInt32

	sx = x0
	sy = y0
	newPath = DoubleInt32{ArrayInt32{sx, sy}}

	for i = 2; i < pathlen; i++ {
		coord = path[i]
		ex = coord[0]
		ey = coord[1]
		line = interpolate(sx, sy, ex, ey)

		blocked = false
		for j = 1; j < len(line); j++ {
			testCoord = line[j]

			if !grid.IsWalkableAt(int(testCoord[0]), int(testCoord[1])) {
				blocked = true
				break
			}
		}

		if blocked {
			lastValidCoord := path[i-1]
			newPath = append(newPath, lastValidCoord)
			sx = lastValidCoord[0]
			sy = lastValidCoord[1]
		}
	}

	newPath = append(newPath, ArrayInt32{x1, y1})

	return newPath
}

/*
 * Compress a path, remove redundant nodes without altering the shape
 * The original path is not modified
	param: path [][]int
	return: [][]int
*/

func compressPath(path DoubleInt32) DoubleInt32 {

	// nothing to compress
	if len(path) < 3 {
		return path
	}

	var (
		compressed = DoubleInt32{}
		sx         = path[0][0] // start X
		sy         = path[0][1] // start Y
		px         = path[1][0] // second point X
		py         = path[1][1] // second point Y
		dx         = px - sx    // direction between the two points
		dy         = py - sy    // direction between the two points
		lx, ly     int32
		ldx, ldy   int32
		sq         int32
		i          int
	)

	// normalize the direction
	sq = int32(math.Ceil(math.Sqrt(float64(dx*dx + dy*dy))))
	dx /= sq
	dy /= sq

	// start the new path
	compressed = append(compressed, []int32{sx, sy})

	for i = 2; i < len(path); i++ {
		// store the last point
		lx = px
		ly = py

		// store the last direction
		ldx = dx
		ldy = dy

		// next point
		px = path[i][0]
		py = path[i][1]

		// next direction
		dx = px - lx
		dy = py - ly

		// normalize
		sq = int32(math.Ceil(math.Sqrt(float64(dx*dx + dy*dy))))
		dx /= sq
		dy /= sq

		// if the direction has changed, store the point
		if dx != ldx || dy != ldy {
			compressed = append(compressed, []int32{lx, ly})
		}
	}

	// store the last point
	compressed = append(compressed, []int32{px, py})
	return compressed
}
