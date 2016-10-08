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
    var _stack = stack.NewStack()
    var node   = self.AsDPNode(n)
    var val    = node.Ints.Peek().(*Vertex).value

    //early exit
    if val < res {
        return
    }

    _stack.Add(n)

    for !_stack.IsEmpty() {
        n = _stack.Pop().(*bst.Node)
        node   = n.Key.(*Node)

        val = node.Ints.Peek().(*Vertex).value

        if !math.IsNaN(val) && val <= res {
            self.NodeSet.Add(n)
        } else {
            _stack.Add(n.Right)
            _stack.Add(n.Left)
        }
    }
}
