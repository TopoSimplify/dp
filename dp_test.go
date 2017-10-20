package dp

import (
	"testing"
	"simplex/offset"
	"github.com/intdxdt/geom"
	"github.com/franela/goblin"
)

func TestDP(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("DP", func() {
		g.It("douglas peucker algorithm", func() {
			var data = []*geom.Point{{0.5, 1.0}, {1.0, 2.0}, {1.0, 0.4}, {2.0, 1.4}, {2.0, 0.8}, {2.5, 1.0},}
			var tree = New(data, offset.MaxOffset)
			tree.Simplify(0.0)
			g.Assert(tree.Simple()).Eql([]int{0, 1, 2, 3, 4, 5})
		})
	})
}

func TestDP2(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("DP2", func() {

		g.It("dp with self intersection", func() {
			var data = []*geom.Point{
				{3.0, 1.6}, {3.0, 2.0}, {2.4, 2.8},
				{0.5, 3.0}, {1.2, 3.2}, {1.4, 2.6}, {2.0, 3.5},
			}

			var tree = New(data, offset.MaxOffset)
			tree.Simplify(0)
			g.Assert(tree.Simple()).Eql(
				[]int{0, 1, 2, 3, 4, 5, 6},
			)
			g.Assert(tree.Simplify(1).Simple()).Eql([]int{0, 3, 6})
			g.Assert(tree.Simplify(3).Simple()).Equal([]int{0, 6})

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
				var data = []*geom.Point{
					{2, 0}, {4, 0}, {4, 1}, {4, 2}, {6, 2}, {8, 2}, {10, 2},
					{10, 1}, {10, 0}, {11, 0}, {12, 0}}

				var tree = New(data, offset.MaxOffset)
				tree.Simplify(0)
				g.Assert(tree.Simple()).Eql([]int{0, 1, 3, 6, 8, 10})
			})
		})

		g.Describe("DP2-0-1-2", func() {
			g.It("dp with empty data", func() {
				var data = []*geom.Point{}
				var tree = New(data, offset.MaxOffset)
				tree.Simplify(0)
				g.Assert(tree.Simple()).Eql([]int{})
			})

			g.It("dp with one coordinate item", func() {
				var data = []*geom.Point{{3.0, 1.6}}
				var tree = New(data, offset.MaxOffset)
				tree.Simplify(0)
				g.Assert(tree.Simple()).Eql([]int{})
				g.Assert(tree.Simplify(1).Simple()).Eql([]int{})
			})

			g.It("dp with two coordinate items", func() {
				var data = []*geom.Point{{3.0, 1.6}, {3.0, 2.0}}
				var tree = New(data, offset.MaxOffset)
				tree.Simplify(0)
				g.Assert(tree.Simple()).Eql([]int{0, 1})
			})
		})
	})
}
