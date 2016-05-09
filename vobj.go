package dp

import (
    . "simplex/struct/item"
    . "simplex/util/math"
    "strconv"
)

type VObj struct {
    index Int
    value float64
}

func (self *VObj) Index() int{
    return int(self.index)
}

func (self *VObj) Value() float64{
    return self.value
}

func NewVObj(i int , v float64) *VObj{
    return &VObj{index : Int(i) , value : v }
}

func (self *VObj) Compare(o Item) int {
    v := o.(*VObj)
    dx := self.value - v.value
    if FloatEqual(dx, 0.0) {
        return 0
    } else if dx < 0 {
        return -1
    }
    return 1
}

func (self VObj) String() string {
    return "{" +
        strconv.Itoa(int(self.index)) +
        ", " +
        strconv.FormatFloat(self.value, 'f', -1, 64) +
        "}"
}

