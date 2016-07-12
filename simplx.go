package dp

import (
    . "simplex/struct/item"
    . "simplex/struct/sset"
)


//Type Simplex
type Simplex struct {
    at *SSet
    rm *SSet
    n  int
}

//New Simplex
func NewSimplex(n int) *Simplex {
    if n < 2 {
        n = 0;
    }
    var self = &Simplex{
        at : NewSSet(),
        rm : NewSSet(),
        n  : n,
    }

    return self.Reset()
}

func (self *Simplex) Add(vals ...int) {
    for _, v := range vals {
        self.UpdateInt(Int(v))
    }
}

func (self *Simplex ) AddSet(set *SSet){
    set.Each(func(o Item) {
        self.UpdateInt(o.(Int))
    })
}

func (self *Simplex ) UpdateInt(i Int) {
    if self.rm.Contains(i) {
        self.at.Add(i)
        self.rm.Remove(i)
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
