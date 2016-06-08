package dp

import (
    . "simplex/struct/item"
    . "simplex/util/math"
    "strconv"
)


type Vertex struct {
    index Int
    value float64
}

func (self *Vertex) Index() Int{
    return self.index
}

func (self *Vertex) Value() float64{
    return self.value
}

func NewVObj(i int , v float64) *Vertex {
    return &Vertex{index : Int(i) , value : v }
}

func (self *Vertex) Compare(o Item) int {
    var v  = o.(*Vertex)
    var dx = self.value - v.value
    if FloatEqual(dx, 0.0) {
        return 0
    } else if dx < 0 {
        return -1
    }
    return 1
}

//vertex to string
func (self Vertex) String() string {
    return "{" +
            strconv.Itoa(int(self.index)) +
        ", " +
            strconv.FormatFloat(self.value, 'f', -1, 64) +
        "}"
}
