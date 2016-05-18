package main

import (
    . "simplex/util/math"
    . "simplex/geom"
    "simplex/dp"
    "fmt"
    "simplex/struct/bst"
    "simplex/struct/item"
    "strconv"
)

var data = []*Point{{0.5, 1.0}, {1.5, 2.0}, {2.5, 1.5}, {3.5, 2.5}, {4.0, 1.5}, {3.0, 1.0} }

func main() {
    //var gstr = 'LINESTRING (520.3891360357894 542.3912033070129, 506.3024618690045 551.4232473315985, 499.8456492240652 555.3948968460392, 492.961552805167 552.5004635914114, 489.3155900796462 547.0315195031302, 494.7910190818659 540.6453203655232, 503.2430235819369 542.0539877822016, 506.3024618690045 551.4232473315985, 505.72509579166825 560.3502151427206, 505.2252456091915 568.0786679640912)'
    //var g = geom.readwkt(gstr)
    var opts = dp.Options{Polyline: data, Threshold: 0}
    var tree = dp.NewDP(opts, true)
    var tree_str = tree.Print()
    fmt.Println(tree_str)

    var o = NewLineString(tree.Coordinates())
    fmt.Println(o.String())
    fmt.Println("gen :", tree.GenInts())
    //tree.Simplify(0.6)
    //var s = _.at(tree.pln, tree.simple.at)
    //o = geom.LineString(s)
    //
    //console.log(o.toString())
    //console.log(s)
    //console.log("\n")
    //
    //var int = tree.int
    //
    //console.log(tree.print(tree.root, _keygen))
    //console.log("\nerror threshold - 0 units\n")
}

func keygen(itm item.Item) string {
    node := itm.(*bst.Node)
    dpnode := node.Key.(*dp.Node)
    ints := dpnode.Ints

    inval := ints.Last().(*dp.Vertex)
    key := dpnode.Key

    var _val = Round(inval.Value(), 1)
    var _int = inval.Index()
    return "{" + strconv.Itoa(_int) + key.String() + "ε:" + fmt.Sprintf("%v", _val) + "}"
}