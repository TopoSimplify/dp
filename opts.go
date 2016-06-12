package dp

import (
    "simplex/struct/item"
    "simplex/struct/rtree"
    . "simplex/relations"
    . "simplex/geom"
)

type Options struct {
    Polyline   []*Point
    Threshold  float64
    Process    func(item.Item)
    Deflection *LineDeflection
    Db         *rtree.RTree
    Relations []Relations
    AvoidSelfIntersection bool
    PreserveComplex  bool
}


func (self *Options) SetPolyline(coords  []*Point)*Options{
    self.Polyline = coords
    return self
}

func (self *Options) SetThreshold(threshold float64)*Options{
    self.Threshold = threshold
    return self
}

func (self *Options) SetProcess(fn  func(item.Item))*Options{
    self.Process = fn
    return self
}

func (self *Options) SetDeflection (deflection  *LineDeflection)*Options{
    self.Deflection = deflection
    return self
}

func (self *Options) SetDb (db  *rtree.RTree)*Options{
    self.Db = db
    return self
}

func (self *Options) SetRelations(relates ...Relations)*Options{
    self.Relations = relates
    return self
}

func (self *Options) SetAvoidSelfIntersection(avoidself  bool)*Options{
    self.AvoidSelfIntersection = avoidself
    return self
}


func (self *Options) SetPreserveComplex(complx  bool)*Options{
    self.PreserveComplex = complx
    return self
}





