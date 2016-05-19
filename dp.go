package dp

import (
    . "simplex/geom"
    . "simplex/struct/item"
    "simplex/struct/bst"
    "simplex/struct/stack"
    "simplex/struct/sset"
    "math"
)

//Type DP
type DP struct {
    *bst.BST
    pln     []*Point
    res     float64
    Simple  *Simplex
    nodeset *sset.SSet
}

//DP constructor
func NewDP(options Options, build bool) *DP {
    var self = &DP{BST: bst.NewBST()}
    var isline, n = self.is_linear_coords(options.Polyline)
    //opts
    self.pln = options.Polyline
    self.res = options.Threshold
    self.nodeset = sset.NewSSet()
    self.Simple = NewSimplex(n)

    fn := options.Process

    if build  && isline {
        self.Build(fn)
    }
    return self
}
//Polyline
func (self *DP) Coordinates() []*Point {
    return self.pln;
}

//Polyline
func (self *DP) is_linear_coords(coords []*Point) (bool, int) {
    n := len(coords)
    if n < 2 {
        n = 0;
    }
    return  n >= 2 , n
}

/*
 description build DP - public use of DP._build
 param [process][Function]
 return {DP}
 */
func (self *DP) Build(process ...func(*bst.Node)) *DP {
    fn := func(*bst.Node) {}
    if len(process) > 0 && process[0] != nil {
        fn = process[0]
    }
    return self.build(fn)
}


//Douglas Peucker BST
//  uses iteration to build tree, state managed with a array stack
//  Note:   int fn_int must return Int type sorted by most interesting
//param [process][Function]   - optional process node callback
//  should return node after process
//  signature : process(node) node{}
//param process
//returns {DP}
func (self *DP)  build(process func(*bst.Node)) *DP {

    var index int
    var val float64

    var range_ = NewInt2D(0, len(self.pln) - 1)

    var n, l, r *bst.Node
    var stack = stack.NewStack()

    var node *Node
    var root = bst.NewNode(NewNode(range_))
    self.BST.Root = root

    //pre stack
    stack.Add(root)

    for !stack.IsEmpty() {
        n = stack.Pop().(*bst.Node)
        process(n)

        node = n.Key.(*Node)
        range_ = node.Key

        node.Ints = self.Offsets(node)
        vobj := node.Ints.Last().(*Vertex)

        index = int(vobj.index)
        val = vobj.value

        if !math.IsNaN(val) && val <= self.res {
            self.Simple.Add(range_[:]...)
        } else {
            //left and right branch
            l = bst.NewNode(
                NewNode(NewInt2D(range_[0], index)),
            )
            r = bst.NewNode(
                NewNode(NewInt2D(index, range_[1])),
            )

            //update pointers
            bst.Ptr(n, l, bst.NewBranch().AsLeft())
            bst.Ptr(n, r, bst.NewBranch().AsRight())

            //node stack
            stack.Add(r)
            stack.Add(l)
        }
    }

    return self
}

//Get all i
func (self *DP) At() []*Point {
    return setvals_coords(self.pln, self.Simple.at)
}

//Get all removed points
func (self *DP) Rm() []*Point {
    return setvals_coords(self.pln, self.Simple.rm)
}
