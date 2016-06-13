package dp

import (
    . "simplex/struct/item"
    "simplex/struct/bst"
    "simplex/struct/stack"
    "math"
)

//Douglas Peucker BST
//  uses iteration to build tree, state managed with a array stack
//  Note:   int fn_int must return Int type sorted by most interesting
//process[Function]   - optional process node callback
//  should return node after process
//  signature : process(node) node{}
//returns *DP
func (self *DP) Build(process ...func(item Item)) *DP {
    procFn := func(item Item) {}
    if len(process) > 0 && process[0] != nil {
        procFn = process[0]
    }
    var offsetter  = self.OffsetFn

    if offsetter == nil {
        offsetter = self
    }


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
        procFn(n.Key)

        node = self.AsDPNode(n)
        range_ = node.Key

        node.Ints = offsetter.Offsets(node)

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