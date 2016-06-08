package dp

import (
    . "simplex/struct/item"
    . "simplex/util/math"
    "simplex/struct/bst"
    "strconv"
    "fmt"
)


//print tree structure as string
//param node{Object} - node
//param key{String|Function} - key attribute
func (self *DP) Print() string {
    return self.BST.Print(keygen)
}
//key generation
func keygen(itm Item) string {
    n := itm.(*bst.Node)
    node := n.Key.(*Node)
    ints := node.Ints
    inval := ints.Peek().(*Vertex)
    key := node.Key

    var _val = Round(inval.Value(), 1)
    var _int = inval.Index().AsInteger()
    return "{" +
        strconv.Itoa(_int) + key.String() +
        "ε:" + fmt.Sprintf("%v", _val) +
        "}"
}

