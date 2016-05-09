package dp

import (
    . "simplex/geom"
    "simplex/struct/item"
    "simplex/struct/bst"
    "simplex/struct/stack"
    "simplex/struct/sset"
    "math"
)



//Type DP
type DP struct {
    *bst.BST
    pln    []*Point
    gen    []int
    res    float64
    simple *Simplex
}

//DP constructor
func NewDP(options Options, build bool) *DP {
    self := &DP{BST: bst.NewBST()}
    //opts
    self.pln = options.Polyline
    self.gen = make([]int, 0)
    self.res = options.Threshold
    self.simple = &Simplex{
        At: make([]int, 0, len(self.pln)),
        Rm: make([]int, 0),
    }
    fn := options.Process
    if build {
        self.Build(fn)
    }
    return self
}

/*
 description  default dp simplification
   returns the leaves of the tree without horizontal or vertical lines
   if callback is not passed
   use callback to modify the behavior of traversing the node
 param node{Object}  - default tree.tree (root)
 param [res]{Number}   - default tree.res
 param [filter] {Function} [Optional] --> func to filter intres vertices
 returns DP
 */
//func (self *DP) simplify(node, res float64, filter) {
//
//  node = is_nil(node) ? self.root : node
//
//  res = is_nil(res) ? self.res : res
//
//  var dpf = is_function(filter) ?
//            filter(self) : Filter(self)
//
//  dpf.filter(node, res)
//
//  while dpf.nodeset.size() {
//    node = dpf.nodeset.shift()
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

/*
 description build DP - public use of DP._build
 param [process][Function]
 return {DP}
 */
func (self *DP) Build(process ...func(*bst.Node)) *DP {
    fn := func(*bst.Node) {}
    if len(process) > 0 && process[0] != nil {
        fn = process[0]
    }
    return self.build(fn)
}


/*
 description Douglas Peucker BST
 *    uses iteration to build tree, state managed with a array stack
 *    Note:   int fn_int must return Int type sorted by most interesting
 param [process][Function]   - optional process node callback
 *    should return node after process
 *    signature : process(node) node{}
 param process
 returns {DP}
 */
func (self *DP)  build(process func(*bst.Node)) *DP {

    var index int
    var val float64

    var idxs = item.NewInt2D(0, len(self.pln) - 1)

    var node, nleft, nright *bst.Node
    var stack = stack.NewStack()
    var intset = sset.NewSSet()

    var dpnode, dpLnode, dpRnode *DPNode
    var root = bst.NewNode(NewDPNode(idxs))
    self.BST.Root = root

    //pre stack
    stack.Add(root)

    for !stack.IsEmpty() {
        node = stack.Pop().(*bst.Node)
        process(node)
        dpnode = node.Key.(*DPNode)
        idxs = dpnode.Key.(*item.Int2D)
        dpnode.Ints = self.Offsets(dpnode)
        vobj := dpnode.Ints.Last().(*VObj)

        index = int(vobj.index)
        val = vobj.value

        if !math.IsNaN(val) && val <= self.res {
            intset.Add(item.Int(idxs[0]))
            intset.Add(item.Int(idxs[1]))
        } else {
            //left and right branch
            dpLnode = NewDPNode(item.NewInt2D(idxs[0], index))
            dpRnode = NewDPNode(item.NewInt2D(index, idxs[1]))

            nleft = bst.NewNode(dpLnode)
            nright = bst.NewNode(dpRnode)

            //update pointers
            bst.Ptr(node, nleft, bst.NewBranch().AsLeft())
            bst.Ptr(node, nright, bst.NewBranch().AsRight())

            //node stack
            stack.Add(nright)
            stack.Add(nleft)
        }
    }

    //gen
    intset.Each(func(itm item.Item) {
        self.gen = append(self.gen, int(itm.(item.Int)))
    })
    return self
}

/*
 description print tree structure as string
 param node{Object} - node
 param key{String|Function} - key attribute
 returns {String|nil}
 */
//func (self *DP) print() {
//  var int = tree.int
//
//  node = is_nil(node) ?
//         tree.root : node
//
//  key = is_nil(key) ?
//        _keygen : key
//
//  return bst.bst.prototype.print.call(tree, node, key)
//
//  func _keygen(node) {
//    var _val = round(int.val(node.int), 3)
//    var _int = int.index(node.int)
//    var _key = node[node._key]
//    return "(" + _int + ", " + _val + " [" + _key + "]" + ")"
//  }
//}
