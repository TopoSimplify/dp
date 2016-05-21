package dp

import (
    . "simplex/struct/sset"
    . "simplex/geom"
    . "simplex/struct/item"
)

func setvals_coords(pln []*Point,  set *SSet)[]*Point{
    coords := make([]*Point, 0)
    set.Each(func(i Item) {
        coords = append(coords,  pln[int(i.(Int))])
    })
    return coords
}

//Get all ints in a set
func setvals_ints(set *SSet) []int {
    at := make([]int, 0);
    set.Each(func(i Item) {
        at = append(at, int(i.(Int)))
    })
    return at
}
