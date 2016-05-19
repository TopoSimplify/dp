package dp

//dp simplification
//  returns the leaves of the tree without horizontal or vertical lines
//  if callback is not passed
//  use callback to modify the behavior of traversing the node
//param node{Object}  - default tree.tree (root)
//param [res]{Number}   - default tree.res
//param [filter] {Function} [Optional] --> func to filter intres vertices
//returns DP
func (self *DP) Simplify(threhold ...float64) *DP {
    var res float64

    if len(threhold) > 0 {
        res = threhold[0]
    } else {
        res = self.res
    }
    var n = self.Root
    var node *Node

    self.Filter(n, res)

    for !self.nodeset.IsEmpty() {
        node = self.nodeset.Shift().(*Node)
        if node == nil {
            continue
        }
        //keep idxs interesting index
        self.Simple.Add(node.Key[:]...)
    }

    return self
}
