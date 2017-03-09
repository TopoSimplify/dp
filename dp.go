package dp

import (
    "simplex/geom"
    "simplex/struct/bst"
    "simplex/struct/item"
    "simplex/struct/sset"
    "simplex/struct/heap"
)

type Offsetter interface {
    Offsets(node *Node) *heap.Heap
}

//Type DP
type DP struct {
    *bst.BST
    Pln     []*geom.Point
    Res     float64
    Simple  *Simplex
    NodeSet *sset.SSet
    Offset  Offsetter
    opts    *Options   //options
    offset  *DPOffsets //default offsetter
}

//DP constructor
func NewDP(options *Options, build bool) *DP {
    var self = &DP{BST: bst.NewBST()}

    self.opts = options
    self.Pln = self.opts.Polyline
    self.Res = self.opts.Threshold
    self.NodeSet = sset.NewSSet()

    self.offset = &DPOffsets{Pln : self.Pln}

    var isline, n = self.is_linear_coords(self.Pln)
    self.Simple = NewSimplex(n)

    fn := options.Process

    if build && isline {
        self.Build(fn)
    }
    return self
}

//Polyline
func (self *DP) Coordinates() []*geom.Point {
    return self.Pln
}

//Polyline
func (self *DP) is_linear_coords(coords []*geom.Point) (bool, int) {
    n := len(coords)
    if n < 2 {
        n = 0
    }
    return n >= 2, n
}

//Get all i
func (self *DP) At() []*geom.Point {
    return setvals_coords(self.Pln, self.Simple.at)
}

//Get all removed points
func (self *DP) Rm() []*geom.Point {
    return setvals_coords(self.Pln, self.Simple.rm)
}

//convert to dp node
func (self *DP) AsDPNode(node *bst.Node) *Node {
    return node.Key.(*Node)
}

//convert to bst node
func (self *DP) AsBSTNode_Item(item item.Item) *bst.Node {
    return item.(*bst.Node)
}

//convert to bst node
func (self *DP) AsBSTNode_Any(item interface{}) *bst.Node {
    return item.(*bst.Node)
}

//convert to dp node from bst node as item interface
func (self *DP) AsDPNode_BSTNode_Item(item item.Item) *Node {
    return self.AsDPNode(self.AsBSTNode_Item(item))
}

//convert to dp node from bst node as item interface
func (self *DP) AsDPNode_BSTNode_Any(any interface{}) *Node {
    return self.AsDPNode(self.AsBSTNode_Any(any))
}

//convert to dp range
func (self *DP) AsDPRange(node *bst.Node) *item.Int2D {
    return self.AsDPNode(node).Key
}
