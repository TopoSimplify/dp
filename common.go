package dp

import (
    . "simplex/geom"
    . "simplex/struct/item"
    "strconv"
)

type Options struct {
    Polyline  []*Point
    Threshold float64
}

//Type Simplex
type Simplex struct{
    At  []int
    Rm  []int
}

//Index range
type Range [2]int
//New Range
func NewRange(i, j int) *Range{
    return &Range{i, j}
}


func (self *Range) Compare(o Item) int {
    v := o.(*Range)
    dx := self - v
    if dx > 0 {
        return 1
    } else if dx < 0 {
        return -1
    }
    return 0 //self - *v
}

func (self *Range) String() string {
    return strconv.Itoa(int(self))
}