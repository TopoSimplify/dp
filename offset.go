package dp

import (
    "simplex/vect"
    . "simplex/struct/slist"
)

// computes euclidean offset distance from dp - archor line
//@param indx{Array} - [i, j] indices
func (self *DP) Offsets(dpnode *DPNode) *SList {
    var dist float64
    var rnge = dpnode.Range
    var N = (rnge[1] - rnge[0]) - 1
    var intlist = NewSList(N)
    var pln = self.pln

    opts := &vect.Options {
        A:self.pln[rnge[0]],
        B:self.pln[rnge[1]],
    }
    anchor := vect.NewVect(opts)

    if N > 0 {
        for i := rnge[0] + 1; i < rnge[1]; i++ {
            dist = anchor.DistanceToPoint(pln[i])
            intlist.Add(NewVObj(i, dist))//store all index,dist
        }
    } else {
        intlist.Add(NewVObj(rnge[1], 0.0))
    }
    return intlist
}
