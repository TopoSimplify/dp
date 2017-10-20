package dp

import (
	"simplex/rng"
	"simplex/node"
	"github.com/intdxdt/deque"
	"github.com/intdxdt/stack"
)

//Douglas Peucker decomposition at a given threshold
func (self *DouglasPeucker) Decompose(threshold float64) *deque.Deque {
	var k int
	var val float64
	var polyline = self.Polyline()
	var hque = deque.NewDeque()

	if polyline == nil {
		return deque.NewDeque()
	}
	var rg = polyline.Range()


	var s = stack.NewStack().Push(rg)

	for !s.IsEmpty() {
		rg = s.Pop().(*rng.Range)
		k, val = self.Score(self, rg)
		if val <= threshold {
			hque.Append(node.New(polyline, rg, hullGeom))
		} else {
			s.Push(
				rng.NewRange(k, rg.J()), // right
				rng.NewRange(rg.I(), k), // left
			)
		}
	}
	return hque
}
