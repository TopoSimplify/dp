package dp

import (
    . "simplex/struct/bst"
    . "simplex/struct/item"
    . "simplex/geom"
    "simplex/struct/slist"
)

type DPNode  struct {
    *Node
    hull *Polygon
    ints slist.SList
}

//Douglas-Peucker Node
func NewDPNode(val Item) *DPNode {
    return &DPNode{NewNode(val), nil}
}
