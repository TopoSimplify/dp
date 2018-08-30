package dp

import (
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/iter"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/sset"
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/lnr"
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/opts"
	"github.com/TopoSimplify/common"
	"github.com/TopoSimplify/decompose"
	"github.com/TopoSimplify/state"
)

//Type DP
type DouglasPeucker struct {
	id        int
	Hulls     []node.Node
	Pln       pln.Polyline
	Meta      map[string]interface{}
	Opts      *opts.Opts
	Score     lnr.ScoreFn
	SimpleSet *sset.SSet
	State     state.State
}

//Creates a new constrained DP Simplification instance
func New(
	id int,
	coordinates geom.Coords,
	options *opts.Opts,
	offsetScore lnr.ScoreFn,
) *DouglasPeucker {
	var instance = DouglasPeucker{
		id:        id,
		Opts:      options,
		Meta:      make(map[string]interface{}, 0),
		SimpleSet: sset.NewSSet(cmp.Int),
		Score:     offsetScore,
	}

	if coordinates.Len() > 1 {
		instance.Pln = pln.CreatePolyline(coordinates)
	}
	return &instance
}

func (self *DouglasPeucker) ScoreRelation(val float64) bool {
	return val <= self.Opts.Threshold
}

func (self *DouglasPeucker) Decompose(id *iter.Igen) []node.Node {
	return decompose.DouglasPeucker(
		id,
		self.Polyline(),
		self.Score,
		self.ScoreRelation,
		common.Geometry,
		self,
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

func (self *DouglasPeucker) Id() int {
	return self.id
}

func (self *DouglasPeucker) Options() *opts.Opts {
	return self.Opts
}

func (self *DouglasPeucker) Coordinates() geom.Coords {
	return self.Pln.Coordinates
}

func (self *DouglasPeucker) Polyline() pln.Polyline {
	return self.Pln
}
