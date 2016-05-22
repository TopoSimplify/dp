package dp

//Simplification at threshold
func (self *DP) Simplify(threhold ...float64) *DP {
    //reset simple sets: at , rm
    self.Simple.Reset()

    var res float64

    if len(threhold) > 0 {
        res = threhold[0]
    } else {
        res = self.res
    }

    var node *Node

    self.Filter(self.Root, res)

    for !(self.nodeset.IsEmpty()) {
        node = self.nodeset.Shift().(*Node)
        //keep idxs interesting index
        self.Simple.Add(node.Key[:]...)
    }

    return self
}
