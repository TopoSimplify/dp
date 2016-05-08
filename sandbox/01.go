package dp
import "simplex/util"
var _ = require("ldsh")
var geom = require("geom")
var dp = require("../")
var data = [
  [0.5, 1.0, 'V0'],
  [1.5, 2.0, 'V1'],
  [2.5, 1.5, 'V2'],
  [3.5, 2.5, 'V3'],
  [4.0, 1.5, 'V4'],
  [3.0, 1.0, 'V5']
]

//var gstr = 'LINESTRING (520.3891360357894 542.3912033070129, 506.3024618690045 551.4232473315985, 499.8456492240652 555.3948968460392, 492.961552805167 552.5004635914114, 489.3155900796462 547.0315195031302, 494.7910190818659 540.6453203655232, 503.2430235819369 542.0539877822016, 506.3024618690045 551.4232473315985, 505.72509579166825 560.3502151427206, 505.2252456091915 568.0786679640912)'
//var g = geom.readwkt(gstr)

var tree = dp.DP({pln: data, res: 0})

var o = geom.LineString(tree.pln)
console.log(o.toString())
console.log('gen :', tree.gen)
tree.simplify(0.6)
var s = _.at(tree.pln, tree.simple.at)
o = geom.LineString(s)

console.log(o.toString())
console.log(s)
console.log("\n")

var int = tree.int

console.log(tree.print(tree.root, _keygen))
console.log("\nerror threshold - 0 units\n")

func _keygen(node) {
  var _val = _.round(int.val(node.int), 1)
  var _int = int.index(node.int)
  var _key = node[node._key]
  return data[_int][2] + '[' + _key + ']' + 'ε:' + _val
}

