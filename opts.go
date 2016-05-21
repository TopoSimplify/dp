package dp

import (
    "simplex/geom"
    "simplex/struct/bst"
)

type Options struct {
    Polyline  []*geom.Point
    Threshold float64
    Process   func(*bst.Node)
    Deflection *LineDeflection
}




