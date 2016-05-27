package dp

import (
    "simplex/struct/stack"
    "simplex/struct/bst"
    "math"
)


//node filter at a given res
//param node
//param res
func (self *DP) Filter(n *bst.Node, res float64) {
    if n == nil {
        return;
    }

    self.NodeSet.Empty()
    var stack  = stack.NewStack()
    var node   = self.AsDPNode(n)
    var val    = node.Ints.Last().(*Vertex).value

    //early exit
    if val < res {
        return
    }

    stack.Add(n)

    for !stack.IsEmpty() {
        n = stack.Pop().(*bst.Node)
        node   = n.Key.(*Node)

        val = node.Ints.Last().(*Vertex).value

        if !math.IsNaN(val) && val <= res {
            self.NodeSet.Add(n)
        } else {
            stack.Add(n.Right)
            stack.Add(n.Left)
        }
    }
}
