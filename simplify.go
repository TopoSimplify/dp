package dp

import (
    "simplex/struct/sset"
    "simplex/struct/item"
)


//Simplification at threshold
func (self *DP) Simplify(opts *Options) *DP {

    self.opts = opts
    self.Simple.Reset()

    var res  = opts.Threshold
    var node *Node

    self.Filter(self.Root, res)

    var intset = sset.NewSSet()

    for !(self.NodeSet.IsEmpty()) {
        node = self.AsDPNode_BSTNode_Item(self.NodeSet.Shift())
        //keep idxs interesting index
        intset.Add(item.Int(node.Key[0]))
        intset.Add(item.Int(node.Key[1]))
    }
    self.Simple.AddSet(intset)

    return self
}

