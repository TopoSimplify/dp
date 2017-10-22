package dp

import (
	"simplex/pln"
	"simplex/rng"
	"simplex/lnr"
	"simplex/node"
	"simplex/opts"
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/sset"
	"github.com/intdxdt/deque"
	"github.com/intdxdt/random"
)

//Type DP
type DouglasPeucker struct {
	Id        string
	Hulls     *deque.Deque
	Pln       *pln.Polyline
	Meta      map[string]interface{}
	Opts       *opts.Opts
	ScoreFn   lnr.ScoreFn
	SimpleSet *sset.SSet
}

//Creates a new constrained DP Simplification instance

func New(coordinates []*geom.Point, options *opts.Opts, offsetScore lnr.ScoreFn) *DouglasPeucker {
	var instance = &DouglasPeucker{
		Id:        random.String(10),
		Opts:      options,
		Hulls:     deque.NewDeque(),
		Meta:      make(map[string]interface{}, 0),
		SimpleSet: sset.NewSSet(cmp.Int),
		ScoreFn:   offsetScore,
	}

	if len(coordinates) > 1 {
		instance.Pln = pln.New(coordinates)
	}
	return instance
}

func (self *DouglasPeucker) Simplify() *DouglasPeucker {
	var hull *node.Node
	var que = self.Decompose()

	self.Hulls.Clear()
	self.SimpleSet.Empty()

	for !que.IsEmpty() {
		hull = que.PopLeft().(*node.Node)
		self.Hulls.Append(hull)
		self.SimpleSet.Extend(hull.Range.I(), hull.Range.J())
	}
	return self
}

func (self *DouglasPeucker) Simple() []int {
	var indices = make([]int, self.SimpleSet.Size())
	self.SimpleSet.ForEach(func(v interface{}, i int) bool {
		indices[i] = v.(int)
		return true
	})
	return indices
}

func (self *DouglasPeucker) Options() *opts.Opts {
	return self.Opts
}

func (self *DouglasPeucker) Coordinates() []*geom.Point {
	return self.Pln.Coordinates
}

func (self *DouglasPeucker) Polyline() *pln.Polyline {
	return self.Pln
}

func (self *DouglasPeucker) Score(pln lnr.Linear, rg *rng.Range) (int, float64) {
	return self.ScoreFn(pln, rg)
}
