package dp

import (
	"testing"
	"simplex/node"
	"simplex/opts"
	"simplex/offset"
	"github.com/intdxdt/geom"
	"github.com/franela/goblin"
)

func TestDecompose(t *testing.T) {
	var g = goblin.Goblin(t)

	g.Describe("hull decomposition", func() {
		g.It("should test decomposition of a line", func() {
			// self.relates = relations(self)
			var wkt = "LINESTRING ( 470 480, 470 450, 490 430, 520 420, 540 440, 560 430, 580 420, 590 410, 630 400, 630 430, 640 460, 630 490, 630 520, 640 540, 660 560, 690 580, 700 600, 730 600, 750 570, 780 560, 790 550, 800 520, 830 500, 840 480, 850 460, 900 440, 920 440, 950 480, 990 480, 1000 520, 1000 570, 990 600, 1010 620, 1060 600 )"
			var coordinates = geom.NewLineStringFromWKT(wkt).Coordinates()
			var options = &opts.Opts{}
			var tree = New(coordinates, options, offset.MaxOffset)

			options.Threshold = 120.0
			var hulls = tree.Decompose()

			g.Assert(hulls.Len()).Equal(4)

			options.Threshold = 150.0
			hulls = tree.Decompose()

			g.Assert(hulls.Len()).Equal(1)
			h := hulls.Get(0).(*node.Node)
			g.Assert(h.Range.AsSlice()).Equal([]int{0, len(coordinates) - 1})
		})
	})
}
