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
    pln         []*Point
    gen         []int
    res         float64
    simple      *Simplex
    nodeset     *sset.SSet
}

//DP constructor
func NewDP(options Options, build bool) *DP {
    self := &DP{BST: bst.NewBST()}
    //opts
    self.pln = options.Polyline
    self.gen = make([]int, 0)
    self.res = options.Threshold
    self.simple = &Simplex{
        At: make([]int, 0, len(self.pln)),
        Rm: make([]int, 0),
    }
    fn := options.Process
    if build {
        self.Build(fn)
    }
    return self
}
//Polyline
func (self *DP) Coordinates() []*Point {
    return self.pln;
}

func (self *DP) GenInts( ) []int{
    return self.gen
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


/*
 description Douglas Peucker BST
   uses iteration to build tree, state managed with a array stack
   Note:   int fn_int must return Int type sorted by most interesting
 param [process][Function]   - optional process node callback
   should return node after process
   signature : process(node) node{}
 param process
 returns {DP}
 */
func (self *DP)  build(process func(*bst.Node)) *DP {

    var index int
    var val float64

    var range_ = NewInt2D(0, len(self.pln) - 1)

    var n, l, r *bst.Node
    var stack = stack.NewStack()
    var intset = sset.NewSSet()

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
            intset.Add(Int(range_[0]))
            intset.Add(Int(range_[1]))
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

    //gen
    intset.Each(func(itm Item) {
        self.gen = append(self.gen, int(itm.(Int)))
    })
    return self
}

