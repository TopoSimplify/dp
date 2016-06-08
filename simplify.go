package dp

import (
    "simplex/struct/sset"
    . "simplex/struct/item"
)


//Simplification at threshold
func (self *DP) Simplify(threhold ...float64) *DP {
    //reset simple sets: at , rm
    self.Simple.Reset()

    var res float64

    if len(threhold) > 0 {
        res = threhold[0]
    } else {
        res = self.Res
    }

    var node *Node

    self.Filter(self.Root, res)

    var intset = sset.NewSSet()

    for !(self.NodeSet.IsEmpty()) {
        node = self.AsDPNode_BSTNode_Item(self.NodeSet.Shift())
        //keep idxs interesting index
        intset.Add(Int(node.Key[0]))
        intset.Add(Int(node.Key[1]))
    }
    self.Simple.AddSet(intset)

    return self
}

