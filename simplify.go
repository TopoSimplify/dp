package dp

/*
 dp simplification
   returns the leaves of the tree without horizontal or vertical lines
   if callback is not passed
   use callback to modify the behavior of traversing the node
 param node{Object}  - default tree.tree (root)
 param [res]{Number}   - default tree.res
 param [filter] {Function} [Optional] --> func to filter intres vertices
 returns DP
 */
//func (self *DP) Simplify(threhold ...float64) {
//    var res float64
//    if len(threhold)> 0 {
//        res = threhold[0]
//    }else {
//        res = self.res
//    }
//  var n = self.Root
//  self.Filter(n, res)
//
//  for !self.nodeset.IsEmpty() {
//    node = self.nodeset.Shift()
//    if !node || is_empty(node[node._key]) {
//      continue
//    }
//    //keep idxs interesting index
//    dpf.intset.appendall(node[node._key])
//  }
//
//  self.gen = is_empty(self.gen) ?
//             idxs(self.len(pln)) : self.gen
//  //at
//  self.simple.at = dpf.intset.values()
//  //rm
//  if self.simple.len(at) {
//    self.simple.rm = difference(self.gen, self.simple.at)
//  }
//  else {
//    self.simple.rm = []
//  }
//  return self
//}