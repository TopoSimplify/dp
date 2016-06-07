package dp

import (
    . "simplex/struct/item"
    "math"
    "simplex/struct/bst"
    "simplex/struct/stack"
)
/*
 description build DP - public use of DP._build
 param [process][Function]
 return {DP}
 */
func (self *DP) Build(process ...func(item Item)) *DP {
    fn := func(item Item) {}
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
func (self *DP)  build(process func(item Item)) *DP {

    var index int
    var val float64

    var range_ *Int2D
    var n, l, r *bst.Node
    var stack = stack.NewStack()

    var node *Node
    var root = bst.NewNode(
        NewNode(0, len(self.Pln) - 1),
    )
    self.BST.Root = root

    //pre stack
    stack.Add(root)

    for !stack.IsEmpty() {
        n = self.AsBSTNode_Any(stack.Pop())
        process(n.Key)

        node = self.AsDPNode(n)
        range_ = node.Key

        node.Ints = self.Offsets(node)
        vobj := node.Ints.Peek().(*Vertex)

        index = int(vobj.index)
        val = vobj.value

        if !math.IsNaN(val) && val <= self.Res {
            self.Simple.Add(range_[:]...)
        } else {
            //left and right branch
            l = bst.NewNode(
                NewNode(range_[0], index),
            )
            r = bst.NewNode(
                NewNode(index, range_[1]),
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
