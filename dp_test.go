package dp

import (
    "testing"
    . "github.com/franela/goblin"
    . "simplex/geom"
    . "simplex/util/math"
    "simplex/struct/item"
    "fmt"
)

func TestDP(t *testing.T) {
    g := Goblin(t)

    g.Describe("DP", func() {
        g.It("douglas peucker algorithm", func() {
            var data = []*Point{
                {0.5, 1.0}, {1.0, 2.0},
                {1.0, 0.4}, {2.0, 1.4},
                {2.0, 0.8}, {2.5, 1.0},
            }
            var tree = NewDP(Options{Polyline: data, Threshold: 0}, true)
            g.Assert(tree.Simple.At()).Eql([]int{0, 1, 2, 3, 4, 5})
            g.Assert(tree.Simplify(0).Simple.At()).Eql([]int{0, 1, 2, 3, 4, 5})
            g.Assert(tree.Simplify().Simple.At()).Eql([]int{0, 1, 2, 3, 4, 5})
            g.Assert(tree.Simplify(0).Simple.Rm()).Eql([]int{})
            g.Assert(tree.At()).Eql(data)
        })
    })
}

func TestDP2(t *testing.T) {
    g := Goblin(t)

    g.Describe("DP2", func() {
        g.It("dp with self intersection", func() {
            var data = []*Point{
                {3.0, 1.6}, {3.0, 2.0}, {2.4, 2.8},
                {0.5, 3.0}, {1.2, 3.2}, {1.4, 2.6}, {2.0, 3.5},
            }
            var tree = NewDP(Options{
                Polyline    : data,
                Threshold   : 0,
                Process     : func(item.Item) {},
            }, true)
            fmt.Println(tree.Print())

            g.Assert(tree.Simple.At()).Eql([]int{0, 1, 2, 3, 4, 5, 6})
            g.Assert(tree.Simple.Rm()).Eql([]int{})
            g.Assert(tree.Simplify(0).Simple.At()).Eql([]int{0, 1, 2, 3, 4, 5, 6})
            g.Assert(tree.Simplify(0).Simple.Rm()).Eql([]int{})
            g.Assert(tree.At()).Eql(tree.Coordinates())
            g.Assert(tree.Rm()).Eql([]*Point{})

            var n = tree.BST.Root
            var node = n.Key.(*Node)
            var vect = node.Ints.Last().(*Vertex)
            var root_key = &item.Int2D{0, 6}

            g.Assert(node.String()).Eql(root_key.String())

            g.Assert(node.Key).Eql(root_key)
            g.Assert(vect.index).Eql(item.Int(3))
            g.Assert(Round(vect.value, 5)).Eql(1.58114)

            n = n.Right
            node = n.Key.(*Node)
            vect = node.Ints.Last().(*Vertex)
            fmt.Println(vect)
            g.Assert(vect.String()).Eql(fmt.Sprintf("{%v, %v}", vect.index, vect.value))

            g.Assert(node.Key).Eql(&item.Int2D{3, 6})
            g.Assert(vect.index).Eql(item.Int(5))
            g.Assert(Round(vect.value, 5)).Eql(0.66408)

            n = tree.BST.Root
            n = n.Left
            node = n.Key.(*Node)
            vect = node.Ints.Last().(*Vertex)

            g.Assert(node.Key).Eql(&item.Int2D{0, 3})
            g.Assert(vect.index).Eql(item.Int(2))
            g.Assert(Round(vect.value, 5)).Eql(0.75385)
            g.Assert(tree.Simplify(1).At()).Eql(
                []*Point{data[0], data[3], data[6]},
            )

            g.Assert(tree.Simplify(3).At()).Eql([]*Point{})

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
                var data = []*Point{
                    {2, 0}, {4, 0}, {4, 1}, {4, 2}, {6, 2}, {8, 2}, {10, 2},
                    {10, 1}, {10, 0}, {11, 0}, {12, 0}}

                fmt.Println(NewLineString(data))

                var tree = NewDP(Options{
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
                var data = []*Point{}
                var tree = NewDP(Options{
                    Polyline    : data,
                    Threshold   : 0,
                    Process     : func(item.Item) {},
                }, true)
                fmt.Println(tree.Print())
                g.Assert(tree.Simple.At()).Eql([]int{})
                g.Assert(tree.Simple.Rm()).Eql([]int{})
            })
            g.It("dp with one coordinate item", func() {
                var data = []*Point{{3.0, 1.6}}
                var tree = NewDP(Options{
                    Polyline    : data,
                    Threshold   : 0,
                    Process     : func(item.Item) {},
                }, true)
                g.Assert(tree.Simple.At()).Eql([]int{})
                g.Assert(tree.Simple.Rm()).Eql([]int{})
                g.Assert(tree.Simplify(1).At()).Eql([]*Point{})
            })
            g.It("dp with two coordinate items", func() {
                var data = []*Point{{3.0, 1.6}, {3.0, 2.0}}
                var tree = NewDP(Options{
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
    g := Goblin(t)

    g.Describe("Line Deflection", func() {
        g.It("tests the straight line deflection of a line", func() {
            var dfln = NewLineDeflection()
            g.Assert(Round(dfln.rad_angle, 2)).Equal(3.1)
            dfln = NewLineDeflection(180.0)
            g.Assert(dfln.rad_angle).Equal(Pi)
            g.Assert(dfln.Deflection()).Eql(Pi)
        })
    })
}
