package dp

import (
	"fmt"
	"testing"
	"simplex/geom"
	"simplex/util/math"
	"simplex/struct/sset"
	"simplex/struct/item"
	"github.com/franela/goblin"
)

func TestVertex(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Vertex", func() {
		g.It("should test Vertex", func() {
			a, b := &Vertex{index:item.Int(50), value: 1.0},
				&Vertex{index:item.Int(10), value:2.0};

			c, d := &Vertex{index:item.Int(1), value:0.4},
				&Vertex{index:item.Int(20), value:0.4};

			e, f := &Vertex{index:item.Int(2), value:8.8},
				&Vertex{index:item.Int(2), value:8.8};

			g.Assert(a.Compare(b)).Eql(-1)
			g.Assert(b.Compare(a)).Eql(1)

			g.Assert(c.Compare(d)).Eql(-1)
			g.Assert(d.Compare(c)).Eql(1)

			g.Assert(e.Compare(f)).Eql(0)
			g.Assert(f.Compare(e)).Eql(0)
		})
	})
}

func TestDP(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("DP", func() {
		g.It("douglas peucker algorithm", func() {
			var data = []*geom.Point{
				{0.5, 1.0}, {1.0, 2.0},
				{1.0, 0.4}, {2.0, 1.4},
				{2.0, 0.8}, {2.5, 1.0},
			}
			var opts = NewOptions()
			opts.Polyline = data
			opts.Threshold = 0

			var tree = NewDP(opts, true)
			g.Assert(tree.Simple.At()).Eql([]int{0, 1, 2, 3, 4, 5})
			g.Assert(tree.Simplify(opts).Simple.At()).Eql(
				[]int{0, 1, 2, 3, 4, 5},
			)
			g.Assert(tree.Simplify(opts.SetThreshold(0)).Simple.At()).Eql(
				[]int{0, 1, 2, 3, 4, 5},
			)
			g.Assert(tree.Simplify(opts).Simple.Rm()).Eql([]int{})
			g.Assert(tree.At()).Eql(data)
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
			var opts = &Options{
				Polyline    : data,
				Threshold   : 0,
				Process     : func(item.Item) {},
			}
			var tree = NewDP(opts, true)
			fmt.Println(tree.Print())

			g.Assert(tree.Simple.At()).Eql(
				[]int{0, 1, 2, 3, 4, 5, 6},
			)
			g.Assert(tree.Simple.Rm()).Eql([]int{})
			g.Assert(tree.Simplify(opts).Simple.At()).Eql(
				[]int{0, 1, 2, 3, 4, 5, 6},
			)
			g.Assert(tree.Simplify(opts).Simple.Rm()).Eql([]int{})
			g.Assert(tree.At()).Eql(tree.Coordinates())
			g.Assert(tree.Rm()).Eql([]*geom.Point{})

			var n = tree.BST.Root
			var node = n.Key.(*Node)
			var vect = node.Ints.Peek().(*Vertex)
			var root_key = &item.Int2D{0, 6}

			g.Assert(node.String()).Eql(root_key.String())

			g.Assert(node.Key).Eql(root_key)
			g.Assert(vect.index).Eql(item.Int(3))
			g.Assert(math.Round(vect.value, 5)).Eql(1.58114)

			n = n.Right
			node = n.Key.(*Node)
			vect = node.Ints.Peek().(*Vertex)

			g.Assert(vect.String()).Eql(fmt.Sprintf("{%v, %v}", vect.index, vect.value))

			g.Assert(node.Key).Eql(&item.Int2D{3, 6})
			g.Assert(vect.index).Eql(item.Int(5))
			g.Assert(math.Round(vect.value, 5)).Eql(0.66408)

			n = tree.BST.Root
			n = n.Left
			node = n.Key.(*Node)
			vect = node.Ints.Peek().(*Vertex)

			g.Assert(node.Key).Eql(&item.Int2D{0, 3})
			g.Assert(vect.index).Eql(item.Int(2))
			g.Assert(math.Round(vect.value, 5)).Eql(0.75385)
			g.Assert(tree.Simplify(opts.SetThreshold(1)).At()).Eql(
				[]*geom.Point{data[0], data[3], data[6]},
			)

			g.Assert(tree.Simplify(opts.SetThreshold(3)).At()).Eql([]*geom.Point{})

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

				fmt.Println(geom.NewLineString(data))

				var tree = NewDP(&Options{
					Polyline    : data,
					Threshold   : 0,
					Process     : func(item.Item) {},
				}, true)
				fmt.Println(tree.Print())
				g.Assert(tree.Simple.At()).Eql([]int{0, 1, 3, 6, 8, 10})
				g.Assert(tree.Simple.Rm()).Eql([]int{2, 4, 5, 7, 9})
			})
		})

		g.Describe("DP2-0-1-2", func() {
			g.It("dp with empty data", func() {
				var data = []*geom.Point{}
				var tree = NewDP(&Options{
					Polyline    : data,
					Threshold   : 0,
					Process     : func(item.Item) {},
				}, true)
				fmt.Println(tree.Print())
				g.Assert(tree.Simple.At()).Eql([]int{})
				g.Assert(tree.Simple.Rm()).Eql([]int{})
			})

			g.It("dp with one coordinate item", func() {
				var data = []*geom.Point{{3.0, 1.6}}
				var opts = &Options{}
				opts.SetThreshold(0,
				).SetPolyline(data,
				).SetDb(nil,
				).SetProcess(func(item.Item) {},
				).SetAvoidSelfIntersection(false,
				).SetPreserveComplex(false,
				).SetRelations(
				).SetDeflection(nil)

				var tree = NewDP(opts, true)
				g.Assert(tree.Simple.At()).Eql([]int{})
				g.Assert(tree.Simple.Rm()).Eql([]int{})
				g.Assert(tree.Simplify(opts.SetThreshold(1)).At()).Eql(
					[]*geom.Point{},
				)
			})

			g.It("dp with two coordinate items", func() {
				var data = []*geom.Point{{3.0, 1.6}, {3.0, 2.0}}
				var tree = NewDP(&Options{
					Polyline    : data,
					Threshold   : 0,
					Process     : func(item.Item) {},
				}, true)
				g.Assert(tree.Simple.At()).Eql([]int{0, 1})
				g.Assert(tree.Simple.Rm()).Eql([]int{})
			})
		})
	})
}

func TestLineDefln(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Line Deflection", func() {
		g.It("tests the straight line deflection of a line", func() {
			var dfln = NewLineDeflection()
			g.Assert(math.Round(dfln.rad_angle, 2)).Equal(3.1)
			dfln = NewLineDeflection(180.0)
			g.Assert(dfln.rad_angle).Equal(math.Pi)
			g.Assert(dfln.Deflection()).Eql(math.Pi)
		})
	})
}

func TestNodeConversion(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Node Conversion", func() {
		g.It("tests node conversion", func() {

			var data = []*geom.Point{
				{3.0, 1.6}, {3.0, 2.0}, {2.4, 2.8},
				{0.5, 3.0}, {1.2, 3.2}, {1.4, 2.6}, {2.0, 3.5},
			}

			var tree = NewDP(&Options{
				Polyline    : data,
				Threshold   : 0,
				Process     : func(item.Item) {},
			}, true)

			var root_key = &item.Int2D{0, 6}

			var node_bst = tree.BST.Root
			var node_item item.Item = tree.BST.Root
			var node_any interface{} = tree.BST.Root

			var n0 = tree.AsDPNode(node_bst)
			var n1 = tree.AsDPNode_BSTNode_Item(node_item)
			var n2 = tree.AsDPNode_BSTNode_Any(node_any)
			var n3 = tree.AsDPRange(node_bst)

			g.Assert(n0.String()).Eql(root_key.String())
			g.Assert(fmt.Sprintf("%T", n0)).Equal("*dp.Node")

			g.Assert(n1.String()).Eql(root_key.String())
			g.Assert(fmt.Sprintf("%T", n1)).Equal("*dp.Node")

			g.Assert(n2.String()).Eql(root_key.String())
			g.Assert(fmt.Sprintf("%T", n2)).Equal("*dp.Node")

			g.Assert(n3.String()).Eql(root_key.String())
			g.Assert(fmt.Sprintf("%T", n3)).Equal("*item.Int2D")

			tree.Simple.Reset()

			var at_list = make([]int, 0)
			var rm_list = make([]int, 0)
			tree.Simple.at.Each(func(o item.Item) {
				at_list = append(at_list, int(o.(item.Int)))
			})
			tree.Simple.rm.Each(func(o item.Item) {
				rm_list = append(rm_list, int(o.(item.Int)))
			})

			g.Assert(at_list).Eql([]int{})
			g.Assert(rm_list).Eql([]int{0, 1, 2, 3, 4, 5, 6, })

			var intset = sset.NewSSet()
			intset.Add(item.Int(0))
			intset.Add(item.Int(6))
			intset.Add(item.Int(2))
			intset.Add(item.Int(4))
			tree.Simple.AddSet(intset)

			at_list = make([]int, 0)
			rm_list = make([]int, 0)
			tree.Simple.at.Each(func(o item.Item) {
				at_list = append(at_list, int(o.(item.Int)))
			})
			tree.Simple.rm.Each(func(o item.Item) {
				rm_list = append(rm_list, int(o.(item.Int)))
			})
			g.Assert(at_list).Eql([]int{0, 2, 4, 6, })
			g.Assert(rm_list).Eql([]int{1, 3, 5, })
		})
	})
}
