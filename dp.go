package dp

import (
	"github.com/TopoSimplify/common"
	"github.com/TopoSimplify/decompose"
	"github.com/TopoSimplify/lnr"
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/offset"
	"github.com/TopoSimplify/opts"
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/state"
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/iter"
	"github.com/intdxdt/sset"
)

//Type DP
type DouglasPeucker struct {
	id          int
	Hulls       []node.Node
	Pln         pln.Polyline
	Meta        map[string]interface{}
	Opts        *opts.Opts
	Score       lnr.ScoreFn
	SquareScore lnr.ScoreFn
	SimpleSet   *sset.SSet
	state       state.State
}

//Creates a new constrained DP Simplification instance
func New(
	id int,
	coordinates geom.Coords,
	options *opts.Opts,
	offsetScore lnr.ScoreFn,
	squareOffsetScore ...lnr.ScoreFn,
) *DouglasPeucker {

	var sqrScore lnr.ScoreFn
	if len(squareOffsetScore) > 0 {
		sqrScore = squareOffsetScore[0]
	}

	var instance = DouglasPeucker{
		id:          id,
		Opts:        options,
		Meta:        make(map[string]interface{}, 0),
		SimpleSet:   sset.NewSSet(cmp.Int),
		Score:       offsetScore,
		SquareScore: sqrScore,
	}

	if coordinates.Len() > 1 {
		instance.Pln = pln.CreatePolyline(coordinates)
	}
	return &instance
}

func (self *DouglasPeucker) ScoreRelation(val float64) bool {
	return val <= self.Opts.Threshold
}

func (self *DouglasPeucker) SquareScoreRelation(val float64) bool {
	return val <= (self.Opts.Threshold * self.Opts.Threshold)
}

func (self *DouglasPeucker) Decompose(id *iter.Igen) []node.Node {
	var score = self.Score
	var relation = self.ScoreRelation
	if self.SquareScore != nil {
		score = self.SquareScore
		relation = self.SquareScoreRelation
	}
	var decomp = offset.EpsilonDecomposition{ScoreFn: score, Relation: relation}
	return decompose.DouglasPeucker(
		id, self.Polyline(), decomp, common.Geometry, self,
	)
}

func (self *DouglasPeucker) Simplify(id *iter.Igen) *DouglasPeucker {
	self.Hulls = self.Decompose(id)
	return self
}

func (self *DouglasPeucker) Simple() []int {
	self.SimpleSet.Empty()
	for i := range self.Hulls {
		self.SimpleSet.Extend(self.Hulls[i].Range.I, self.Hulls[i].Range.J)
	}
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

func (self *DouglasPeucker) State() *state.State {
	return &self.state
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
