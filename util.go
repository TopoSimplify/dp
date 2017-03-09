package dp

import (
    "simplex/geom"
    "simplex/struct/sset"
    "simplex/struct/item"
)

func setvals_coords(pln []*geom.Point,  set *sset.SSet)[]*geom.Point{
    coords := make([]*geom.Point, 0)
    set.Each(func(i item.Item) {
        coords = append(coords,  pln[int(i.(item.Int))])
    })
    return coords
}

//Get all ints in a set
func setvals_ints(set *sset.SSet) []int {
    at := make([]int, 0);
    set.Each(func(i item.Item) {
        at = append(at, int(i.(item.Int)))
    })
    return at
}
