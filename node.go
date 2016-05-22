package dp

import (
    . "simplex/struct/item"
    "simplex/geom"
    "simplex/struct/slist"
)

type Node  struct {
    Key  *Int2D
    Hull *geom.Polygon
    Ints *slist.SList
}

func NewNode(i, j int ) *Node {
    return &Node{
        Key  : NewInt2D(i , j),
        Hull : nil,
        Ints : nil,
    }
}

func (self *Node ) Compare(other  Item) int {
    dpother := other.(*Node)
    return self.Key.Compare(dpother.Key)
}

func (self *Node) String() string {
    return self.Key.String()
}