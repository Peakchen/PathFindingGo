package core

/*
	by stefan 2572915286@qq.com
	from https://github.com/qiao/PathFinding.js
*/

type DiagonalMovement int

const (
	Always              DiagonalMovement = 1
	Never               DiagonalMovement = 2
	IfAtMostOneObstacle DiagonalMovement = 3
	OnlyWhenNoObstacles DiagonalMovement = 4
)
