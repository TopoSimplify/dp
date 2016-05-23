package dp

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

    for !(self.NodeSet.IsEmpty()) {
        node = self.NodeSet.Shift().(*Node)
        //keep idxs interesting index
        self.Simple.Add(node.Key[:]...)
    }

    return self
}
