package dp

import (
    . "simplex/geom"
    . "simplex/struct/item"
    "simplex/struct/bst"
    "simplex/struct/sset"
)

type Options struct {
    Polyline  []*Point
    Threshold float64
    Process   func(*bst.Node)
}

//Type Simplex
type Simplex struct {
    at *sset.SSet
    rm *sset.SSet
}

func NewSimplex(n int) *Simplex {
    var self = &Simplex{
        at : sset.NewSSet(),
        rm : sset.NewSSet(),
    }
    for i := 0; i < n; i++ {
        self.rm.Add(Int(i))
    }
    return self
}

func (self *Simplex) Add(vals ...int) {
    var i Int
    for _, v := range vals {
        i = Int(v)
        if self.rm.Contains(i) {
            self.at.Add(i)
            self.rm.Remove(i)
        }
    }
}
//Get all ints at interesting indices
func (self *Simplex) At() []int {
    return setvals_ints(self.at)
}

//Get all removed set
func (self *Simplex) Rm() []int {
    return setvals_ints(self.rm)
}





