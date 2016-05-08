package dp
import "simplex/util"
var _ = require("ldsh")
var geom = require("geom")
var dp = require("../")
var test = require("tape")
test("douglas peucker algorithm", func (t) {
  var data = [
    [0.5, 1.0, "V0"],
    [1.0, 2.0, "V1"],
    [1.0, 0.4, "V2"],
    [2.0, 1.4, "V3"],
    [2.0, 0.8, "V4"],
    [2.5, 1.0, "V5"]
  ]

  var tree = new dp.DP({pln: data, res: 0})
  t.Assert(tree.gen, [0, 1, 2, 3, 4, 5])
  var obj = tree.simplify(0)
  var simplx = _.at(data, tree.simple.at)
  t.Assert(simplx, data)
  t.Assert(obj==tree, true)
  console.log(tree.print())
  t.end()
})

test("dp with self-intersect", func (t) {
  t.plan(9)
  var data = [
    [3.0, 1.6], [3.0, 2.0], [2.4, 2.8],
    [0.5, 3.0], [1.2, 3.2], [1.4, 2.6], [2.0, 3.5]
  ]

  var g = geom.LineString(data)
  console.log(g.toString())

  var node
  var tree = new dp.DP({pln: data, res: 0})
  t.Assert(tree.gen, [0, 1, 2, 3, 4, 5, 6])

  tree.simplify(0)
  var simplx = _.at(data, tree.simple.at)
  t.Assert(simplx, data)

  node = tree.root//root

  t.Assert(node[tree.key], [0, 6])
  t.Assert(tree.int.index(node.int), 3)
  t.Assert(_.round(tree.int.val(node.int), 5), 1.58114)

  node = tree.root.right//root.right
  t.Assert(node[tree.key], [3, 6])
  t.Assert(tree.int.index(node.int), 5)
  t.Assert(_.round(tree.int.val(node.int), 5), 0.66408)

  //simplify at 1
  tree.simplify(1)
  simplx = _.at(data, tree.simple.at)
  t.Assert(simplx, [data[0], data[3], data[6]])
  console.log(tree.print())
})
