package dp

import (
    . "simplex/struct/item"
    "simplex/geom"
    "simplex/struct/slist"
)

type DPNode  struct {
    Key  Item
    Hull *geom.Polygon
    Ints *slist.SList
}

func NewDPNode(item Item) *DPNode {
    return &DPNode{
        Key  : item,
        Hull : nil,
        Ints : nil,
    }
}

func (self *DPNode ) Compare(other  Item) int {
    dpother := other.(*DPNode)
    return self.Key.Compare(dpother.Key)
}

func (self *DPNode) String() string {
    return self.Key.String()
}

