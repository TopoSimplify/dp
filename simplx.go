package dp

import (
    . "simplex/struct/item"
    "simplex/struct/sset"
)


//Type Simplex
type Simplex struct {
    at *sset.SSet
    rm *sset.SSet
    n  int
}

//New Simplex
func NewSimplex(n int) *Simplex {
    if n < 2 {
        n = 0;
    }
    var self = &Simplex{
        at : sset.NewSSet(),
        rm : sset.NewSSet(),
        n  : n,
    }

    return self.Reset()
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


//Reset at and rm indices
func (self *Simplex) Reset() *Simplex{
    self.at.Empty()
    self.rm.Empty()
    for i := 0; i < self.n ; i++ {
        self.rm.Add(Int(i))
    }
    return self
}

//Get all ints at interesting indices
func (self *Simplex) At() []int {
    return setvals_ints(self.at)
}

//Get all removed set
func (self *Simplex) Rm() []int {
    return setvals_ints(self.rm)
}

