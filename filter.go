package dp

import (
    "simplex/struct/sset"
    "simplex/struct/stack"
    "simplex/struct/bst"
)

//import (
//  "simplex/util"
//  "simplex/struct/stack"
//  "simplex/struct/sset"
//)


//node filter at a given res
//param node
//param res
func (self *DP) Filter(n *bst.Node, res float64) {

    self.nodeset.Empty()
    var stack  = stack.NewStack()
    var node   = n.Key.(*Node)
    var val    = node.Ints.Last().(*Vertex).value

    //early exit
    if val < res {
        return
    }

    stack.Add(n)

    for !stack.IsEmpty() {
        n = stack.Pop().(*bst.Node)
        node   = n.Key.(*Node)
        if node == nil {
            continue
        }

        val = node.Ints.Last().(*Vertex).value

        if !(val == nil) && val <= res {
            self.nodeset.Add(node)
        } else {
            stack.Add(n.Right)
            stack.Add(n.Left)
        }
    }
}

