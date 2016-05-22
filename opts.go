package dp

import (
    "simplex/geom"
    "simplex/struct/item"
)

type Options struct {
    Polyline  []*geom.Point
    Threshold float64
    Process   func(item.Item)
    Deflection *LineDeflection
}




