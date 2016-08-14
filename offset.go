package dp

import (
    . "simplex/struct/heap"
    . "simplex/geom"
    "simplex/cart2d"
)
//DP Offsets - Minimum Distance Offsets
type DPOffsets struct {
    Pln []*Point
}

//euclidean offset distance from dp - archor line
//@param indx{Array} - [i, j] indices
func (self *DPOffsets) Offsets(node *Node) *Heap {
    var indx = node.Key
    var dist float64
    var N = (indx[1] - indx[0]) - 1

    var ints = NewHeap(NewHeapType().AsMax())
    if N > 0 {
        for i := indx[0] + 1; i < indx[1]; i++ {
            dist = cart2d.DistanceToPoint(
                self.Pln[indx[0]],
                self.Pln[indx[1]],
                self.Pln[i],
            )
            ints.Push(NewVObj(i, dist))//store all index,dist
        }
    } else {
        ints.Push(NewVObj(indx[1], 0.0))
    }
    return ints
}


