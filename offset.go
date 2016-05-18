package dp

import (
    "simplex/vect"
    . "simplex/struct/slist"
)

//computes euclidean offset distance from dp - archor line
//@param indx{Array} - [i, j] indices
func (self *DP) Offsets(node *Node) *SList {
    var indx = node.Key
    var dist float64
    var N = (indx[1] - indx[0]) - 1
    var intlist = NewSList(N)
    var pln = self.pln

    opts := &vect.Options{
        A:self.pln[indx[0]],
        B:self.pln[indx[1]],
    }
    anchor := vect.NewVect(opts)

    if N > 0 {
        for i := indx[0] + 1; i < indx[1]; i++ {
            dist = anchor.DistanceToPoint(pln[i])
            intlist.Add(NewVObj(i, dist))//store all index,dist
        }
    } else {
        intlist.Add(NewVObj(indx[1], 0.0))
    }
    return intlist
}
