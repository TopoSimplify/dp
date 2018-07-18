package dp

import (
	"time"
	"testing"
	"github.com/intdxdt/geom"
	"github.com/franela/goblin"
	"github.com/TopoSimplify/opts"
	"github.com/TopoSimplify/offset"
)

func TestDP(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("DP", func() {
		g.It("douglas peucker algorithm", func() {
			g.Timeout(1 * time.Hour)

			var data = []geom.Point{{0.5, 1.0}, {1.0, 2.0}, {1.0, 0.4}, {2.0, 1.4}, {2.0, 0.8}, {2.5, 1.0},}
			var options = &opts.Opts{}
			var tree = New(data, options, offset.MaxOffset)
			tree.Simplify()
			g.Assert(tree.Simple()).Eql([]int{0, 1, 2, 3, 4, 5})
		})
	})
}

func TestDP2(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("DP2", func() {
		g.It("dp with self intersection", func() {
			g.Timeout(1 * time.Hour)

			var data = []geom.Point{
				{3.0, 1.6}, {3.0, 2.0}, {2.4, 2.8},
				{0.5, 3.0}, {1.2, 3.2}, {1.4, 2.6}, {2.0, 3.5},
			}
			var options = &opts.Opts{Threshold: 0}
			var tree = New(data, options, offset.MaxOffset)
			tree.Simplify()
			g.Assert(tree.Simple()).Eql([]int{0, 1, 2, 3, 4, 5, 6})
			options.Threshold = 1
			g.Assert(tree.Simplify().Simple()).Eql([]int{0, 3, 6})

			g.Assert(tree.Options()).Equal(tree.Opts)
			g.Assert(tree.NodeQueue()).Equal(tree.Hulls)
			g.Assert(tree.Coordinates()).Equal(data)
			g.Assert(tree.Id()).Equal(tree.id)
			g.Assert(NodeGeometry([]geom.Point{data[0]}).(geom.Point)).Equal(data[0])

			options.Threshold = 3
			g.Assert(tree.Simplify().Simple()).Equal([]int{0, 6})

		})

		g.Describe("Horizontal-vertical", func() {
			g.It("horz-vert", func() {
				/*
						  (3).....(4)....(5)....(6)
						   |                     |
						  (2)                   (7)
						   |                     |
				   (0)....(1)                   (8)....(9)...(10)
				 */
				var data = []geom.Point{
					{2, 0}, {4, 0}, {4, 1}, {4, 2}, {6, 2}, {8, 2}, {10, 2},
					{10, 1}, {10, 0}, {11, 0}, {12, 0}}
				var options = &opts.Opts{}
				var tree = New(data, options, offset.MaxOffset)
				tree.Simplify()
				g.Assert(tree.Simple()).Eql([]int{0, 1, 3, 6, 8, 10})
			})
		})

		g.Describe("DP2-0-1-2", func() {
			g.It("dp with empty data", func() {
				var data = []geom.Point{}
				var options = &opts.Opts{}
				var tree = New(data, options, offset.MaxOffset)
				tree.Simplify()
				g.Assert(tree.Simple()).Eql([]int{})
			})

			g.It("dp with one coordinate item", func() {
				var data = []geom.Point{{3.0, 1.6}}
				var options = &opts.Opts{}
				var tree = New(data, options, offset.MaxOffset)
				tree.Simplify()
				g.Assert(tree.Simple()).Eql([]int{})
				options.Threshold = 1
				g.Assert(tree.Simplify().Simple()).Eql([]int{})
			})

			g.It("dp with two coordinate items", func() {
				var data = []geom.Point{{3.0, 1.6}, {3.0, 2.0}}
				var options = &opts.Opts{}
				var tree = New(data, options, offset.MaxOffset)
				tree.Simplify()
				g.Assert(tree.Simple()).Eql([]int{0, 1})
			})
		})
	})
}
