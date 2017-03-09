package dp

import (
	"math"
	"simplex/geom"
	"simplex/relations"
	"simplex/struct/item"
	"simplex/struct/rtree"
)

type Options struct {
	Polyline              []*geom.Point
	Threshold             float64
	MinimumDist           float64
	Process               func(item.Item)
	Deflection            *LineDeflection
	Db                    *rtree.RTree
	Relations             []relations.Relations
	AvoidSelfIntersection bool
	PreserveComplex       bool
}

func NewOptions() *Options {
	return &Options{MinimumDist: math.MaxFloat64}
}

func (self *Options) SetPolyline(coords []*geom.Point) *Options {
	self.Polyline = coords
	return self
}

func (self *Options) SetThreshold(threshold float64) *Options {
	self.Threshold = threshold
	return self
}

func (self *Options) SetProcess(fn func(item.Item)) *Options {
	self.Process = fn
	return self
}

func (self *Options) SetDeflection(deflection *LineDeflection) *Options {
	self.Deflection = deflection
	return self
}

func (self *Options) SetDb(db *rtree.RTree) *Options {
	self.Db = db
	return self
}

func (self *Options) SetRelations(relates ...relations.Relations) *Options {
	self.Relations = relates
	return self
}

func (self *Options) SetAvoidSelfIntersection(avoidself bool) *Options {
	self.AvoidSelfIntersection = avoidself
	return self
}

func (self *Options) SetPreserveComplex(complx bool) *Options {
	self.PreserveComplex = complx
	return self
}
