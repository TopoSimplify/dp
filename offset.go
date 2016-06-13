package dp

import (
    "simplex/vect"
    . "simplex/struct/heap"
    . "simplex/geom"
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

    opts := &vect.Options{
        A : self.Pln[indx[0]],
        B : self.Pln[indx[1]],
    }
    anchor := vect.NewVect(opts)

    if N > 0 {
        for i := indx[0] + 1; i < indx[1]; i++ {
            dist = anchor.DistanceToPoint(self.Pln[i])
            ints.Push(NewVObj(i, dist))//store all index,dist
        }
    } else {
        ints.Push(NewVObj(indx[1], 0.0))
    }
    return ints
}


