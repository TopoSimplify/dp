package dp

import (
    "simplex/geom"
    "simplex/struct/item"
    "simplex/struct/rtree"
    "simplex/relations"
)

type Options struct {
    Polyline   []*geom.Point
    Threshold  float64
    Process    func(item.Item)
    Deflection *LineDeflection
    Db         *rtree.RTree
    Relations []relations.Relations
    AvoidSelfIntersection bool
    PreserveComplex  bool
}




