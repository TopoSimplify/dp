package dp

import (
    . "simplex/struct/item"
    "simplex/geom"
    "simplex/struct/slist"
)

type DPNode  struct {
    Range *Int2D
    Hull  *geom.Polygon
    Ints  *slist.SList
}

func NewDPNode(rnge *Int2D) *DPNode {
    return &DPNode{
        Range   : rnge,
        Hull    : nil,
        Ints    : nil,
    }
}

func AsDPNode(item Item) *DPNode {
    return item.(*DPNode)
}

func (self *DPNode ) Compare(other  Item) int {
    dpother := other.(*DPNode)
    return self.Range.Compare(dpother.Range)
}

func (self *DPNode) String() string {
    return self.Range.String()
}

