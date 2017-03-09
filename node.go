package dp

import (
	"simplex/geom"
	"simplex/struct/item"
	"simplex/struct/heap"
)

type Node struct {
	Key  *item.Int2D
	Hull *geom.Polygon
	Ints *heap.Heap
}

func NewNode(i, j int) *Node {
	return &Node{
		Key:  item.NewInt2D(i, j),
		Hull: nil,
		Ints: nil,
	}
}

func (self *Node) Compare(other item.Item) int {
	dpother := other.(*Node)
	return self.Key.Compare(dpother.Key)
}

func (self *Node) String() string {
	return self.Key.String()
}
