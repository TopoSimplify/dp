package dp

import "github.com/intdxdt/geom"

//hull geom
func NodeGeometry(coordinates []geom.Point) geom.Geometry {
	var g geom.Geometry
	if len(coordinates) > 2 {
		g = geom.NewPolygon(coordinates)
	} else if len(coordinates) == 2 {
		g = geom.NewLineString(coordinates)
	} else {
		g = coordinates[0]
	}
	return g
}
