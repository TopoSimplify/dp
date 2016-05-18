package dp

import (
    . "simplex/geom"
    "simplex/struct/bst"
)


type Options struct {
    Polyline  []*Point
    Threshold float64
    Process  func (*bst.Node)
}

//Type Simplex
type Simplex struct{
    At  []int
    Rm  []int
}