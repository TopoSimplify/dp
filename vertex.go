package dp

import (
	"strconv"
	"simplex/util/math"
	"simplex/struct/item"
)

type Vertex struct {
	index item.Int
	value float64
}

func (self *Vertex) Index() item.Int {
	return self.index
}

func (self *Vertex) Value() float64 {
	return self.value
}

func NewVObj(i int, v float64) *Vertex {
	return &Vertex{index: item.Int(i), value: v }
}

func (self *Vertex) Compare(o item.Item) int {
	var v = o.(*Vertex)
	var dx = self.value - v.value
	if math.FloatEqual(dx, 0.0) {
		//compare by index
		if self.index > v.index {
			return 1
		} else if self.index < v.index {
			return -1
		}
		return 0
	} else if dx < 0 {
		return -1
	}
	return 1
}

//vertex to string
func (self Vertex) String() string {
	return "{" + strconv.Itoa(int(self.index)) + ", " +
		strconv.FormatFloat(self.value, 'f', -1, 64) + "}"
}
