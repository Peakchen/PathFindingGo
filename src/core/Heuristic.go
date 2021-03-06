package core

import "math"

/*
	by stefan 2572915286@qq.com
	Based upon  https://github.com/qiao/PathFinding.js
*/

/**
 * Manhattan distance.
 * @param {number} dx - Difference in x.
 * @param {number} dy - Difference in y.
 * @return {number} dx + dy
 */
func Manhattan(dx, dy int32) int32 {
	return dx + dy
}

/**
 * Euclidean distance.
 * @param {number} dx - Difference in x.
 * @param {number} dy - Difference in y.
 * @return {number} sqrt(dx * dx + dy * dy)
 */
func Euclidean(dx, dy int32) int32 {
	return int32(math.Ceil(math.Sqrt(float64(dx*dx + dy*dy))))
}

/**
 * Octile distance.
 * @param {number} dx - Difference in x.
 * @param {number} dy - Difference in y.
 * @return {number} sqrt(dx * dx + dy * dy) for grids
 */
func Octile(dx, dy int32) int32 {
	var F = SQRT2 - float64(1)
	if dx < dy {
		return int32(math.Ceil((F*float64(dx) + float64(dy))))
	}
	return int32(math.Ceil((F*float64(dy) + float64(dx))))
}

/**
 * Chebyshev distance.
 * @param {number} dx - Difference in x.
 * @param {number} dy - Difference in y.
 * @return {number} max(dx, dy)
 */
func Chebyshev(dx, dy int32) int32 {
	return int32(math.Max(float64(dx), float64(dy)))
}
