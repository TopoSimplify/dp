package dp

import (
	"simplex/pln"
	"simplex/rng"
	"simplex/lnr"
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/sset"
	"github.com/intdxdt/deque"
	"github.com/intdxdt/random"
	"simplex/node"
)

//Type DP
type DouglasPeucker struct {
	Id    string
	Hulls *deque.Deque
	Pln   *pln.Polyline
	Meta  map[string]interface{}

	score  lnr.ScoreFn
	simple *sset.SSet
}

//Creates a new constrained DP Simplification instance

func New(coordinates []*geom.Point, offsetScore lnr.ScoreFn) *DouglasPeucker {
	var instance = &DouglasPeucker{
		Id:     random.String(10),
		Hulls:  deque.NewDeque(),
		Meta:   make(map[string]interface{}, 0),
		simple: sset.NewSSet(cmp.Int),
		score:  offsetScore,
	}
	if len(coordinates) > 1 {
		instance.Pln = pln.New(coordinates)
	}
	return instance
}

func (self *DouglasPeucker) Simplify(threshold float64) *DouglasPeucker {
	var hull *node.Node
	var que = self.Decompose(threshold)

	self.Hulls.Clear()
	self.simple.Empty()

	for !que.IsEmpty() {
		hull = que.PopLeft().(*node.Node)
		self.Hulls.Append(hull)
		self.simple.Extend(hull.Range.I(), hull.Range.J())
	}
	return self
}

func (self *DouglasPeucker) Simple() []int {
	var indices = make([]int, self.simple.Size())
	self.simple.ForEach(func(v interface{}, i int) bool {
		indices[i] = v.(int)
		return true
	})
	return indices
}

func (self *DouglasPeucker) Coordinates() []*geom.Point {
	return self.Pln.Coordinates
}

func (self *DouglasPeucker) Polyline() *pln.Polyline {
	return self.Pln
}

func (self *DouglasPeucker) Score(pln lnr.Linear, rg *rng.Range) (int, float64) {
	return self.score(pln, rg)
}
