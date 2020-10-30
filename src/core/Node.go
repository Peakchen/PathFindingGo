package core

/*
	by stefan 2572915286@qq.com
	from https://github.com/qiao/PathFinding.js
*/

type DoubleNode [][]*TNode
type ArrayNode []*TNode

type TNode struct {
	x        int32
	y        int32
	walkable bool
	parent   *TNode
}

func Node(x int32, y int32, walkable bool) *TNode {
	return &TNode{
		x:        x,
		y:        y,
		walkable: walkable,
		parent:   nil,
	}
}
