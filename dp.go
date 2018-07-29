package dp

import (
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/sset"
	"github.com/intdxdt/random"
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/lnr"
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/opts"
	"github.com/TopoSimplify/decompose"
	"github.com/intdxdt/iter"
)

//Type DP
type DouglasPeucker struct {
	id        string
	Hulls     []node.Node
	Pln       *pln.Polyline
	Meta      map[string]interface{}
	Opts      *opts.Opts
	Score     lnr.ScoreFn
	SimpleSet *sset.SSet
}

//Creates a new constrained DP Simplification instance
func New(coordinates []geom.Point, options *opts.Opts, offsetScore lnr.ScoreFn) *DouglasPeucker {
	var instance = &DouglasPeucker{
		id:        random.String(10),
		Opts:      options,
		Hulls:     []node.Node{},
		Meta:      make(map[string]interface{}, 0),
		SimpleSet: sset.NewSSet(cmp.Int),
		Score:     offsetScore,
	}

	if len(coordinates) > 1 {
		instance.Pln = pln.New(coordinates)
	}
	return instance
}

func (self *DouglasPeucker) ScoreRelation(val float64) bool {
	return val <= self.Opts.Threshold
}

func (self *DouglasPeucker) Decompose(id *iter.Igen) []node.Node {
	return decompose.DouglasPeucker(
		id, self.Polyline(), self.Score,
		self.ScoreRelation, NodeGeometry,
	)
}

func (self *DouglasPeucker) Simplify(id *iter.Igen) *DouglasPeucker {
	self.SimpleSet.Empty()
	self.Hulls = self.Decompose(id)

	for i := range self.Hulls {
		self.SimpleSet.Extend(self.Hulls[i].Range.I, self.Hulls[i].Range.J)
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

func (self *DouglasPeucker) Id() string {
	return self.id
}

func (self *DouglasPeucker) Options() *opts.Opts {
	return self.Opts
}

func (self *DouglasPeucker) Coordinates() []geom.Point {
	return self.Pln.Coordinates
}

func (self *DouglasPeucker) Polyline() *pln.Polyline {
	return self.Pln
}

