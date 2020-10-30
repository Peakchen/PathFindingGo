package core

import "math"

/*
	by stefan 2572915286@qq.com
	from https://github.com/qiao/PathFinding.js
*/

/**
 * Manhattan distance.
 * @param {number} dx - Difference in x.
 * @param {number} dy - Difference in y.
 * @return {number} dx + dy
 */
func manhattan(dx, dy int32) int32 {
	return dx + dy
}

/**
 * Euclidean distance.
 * @param {number} dx - Difference in x.
 * @param {number} dy - Difference in y.
 * @return {number} sqrt(dx * dx + dy * dy)
 */
func euclidean(dx, dy int32) int32 {
	return int32(math.Ceil(math.Sqrt(float64(dx*dx + dy*dy))))
}

/**
 * Octile distance.
 * @param {number} dx - Difference in x.
 * @param {number} dy - Difference in y.
 * @return {number} sqrt(dx * dx + dy * dy) for grids
 */
func octile(dx, dy int32) float64 {
	var F = SQRT2 - float64(1)
	if dx < dy {
		return (F*float64(dx) + float64(dy))
	}
	return (F*float64(dy) + float64(dx))
}

/**
 * Chebyshev distance.
 * @param {number} dx - Difference in x.
 * @param {number} dy - Difference in y.
 * @return {number} max(dx, dy)
 */
func chebyshev(dx, dy int32) int32 {
	return int32(math.Max(float64(dx), float64(dy)))
}
