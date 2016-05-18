package dp

//import (
//  "simplex/util"
//  "simplex/struct/stack"
//  "simplex/struct/sset"
//)



 //node filter at a given res
 //param node
 //param res

//func(self *DP) filter(node *DPNode, res float64) {
//
//  var intset  = sset.NewSSet()
//  var nodeset = sset.NewSSet()
//  var stack   = stack.NewStack()
//  var vobj    =  node.ints.Last().(*DPNode)
//  var val     = int.val(node.int)
//
//  //early exit
//  if val < res {
//    return self.nodeset
//  }
//
//  stack.append(node)
//
//  while !stack.isempty() {
//    node = stack.pop()
//    if _.is_nil(node) {
//      continue
//    }
//    val = int.val(node.int)
//    if !(val == nil) && val <= res {
//      self.nodeset.append(node)
//    }
//    else {
//      stack.append(node.right)
//      stack.append(node.left)
//    }
//  }
//}

