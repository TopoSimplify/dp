package dp

import (
    "simplex/geom"
    "simplex/struct/item"
    "simplex/struct/rtree"
    "simplex/constrelate"
)

type Options struct {
    Polyline   []*geom.Point
    Threshold  float64
    Process    func(item.Item)
    Deflection *LineDeflection
    Db         *rtree.RTree
    Constraints []*constrelate.Constraint
}




