package dp

import "simplex/util/math"

const DefaultDefln = 177.6169164905552 // 3.1 rad

//Line deflection type
type LineDeflection struct {
	rad_angle float64
}

//Creates new line deflection
func NewLineDeflection(deg ...float64) *LineDeflection {
	var d float64 = DefaultDefln
	if len(deg) > 0 {
		d = deg[0]
	}
	return &LineDeflection{
		math.Deg2rad(d),
	}
}

//Deflection in Radians
func (self *LineDeflection) Deflection() float64 {
	return self.rad_angle
}
