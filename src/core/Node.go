package core

/*
	by stefan 2572915286@qq.com
	from https://github.com/qiao/PathFinding.js
*/

type DoubleNode [][]*TNode
type ArrayNode []*TNode

type TNode struct {
	X        int32
	Y        int32
	Walkable bool
	Parent   *TNode
}

func Node(x int32, y int32, Walkable bool) *TNode {
	return &TNode{
		X:        x,
		Y:        y,
		Walkable: Walkable,
		Parent:   nil,
	}
}

func (this *TNode) IsEqual(node *TNode) bool {
	return this.X == node.X && this.Y == node.Y && this.Walkable == node.Walkable
}
