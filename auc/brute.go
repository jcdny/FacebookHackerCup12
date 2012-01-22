// Copyright © 2012 Jeffrey Davis <jeff.davis@gmail.com>
// Use of this code is governed by the GPL version 2 or later.
// See the file LICENSE for details.

package main

import (
	"fmt"
)

func ComputeCaseBrute(a *Auction) (nt, nb int64) {
	ivec := make([]Item, 0, a.N)
	pi := a.P1
	wi := a.W1
	for i := int64(0); i < a.N; i++ {
		ivec = append(ivec, Item{P: int(pi), W: int(wi)})
		pi = ((a.A*pi+a.B)%a.M + 1)
		wi = ((a.C*wi+a.D)%a.K + 1)
	}
	nt, nb = Brute(a.Case, ivec)
	return
}

func Brute(ic int, ivec []Item) (nt, nb int64) {
	for j, ia := range ivec {
		xt := 0
		xb := 0
		for k, ib := range ivec {
			if k == j {
				continue
			}
			if (ia.P < ib.P && ia.W <= ib.W) || (ia.P <= ib.P && ia.W < ib.W) {
				if xt == 0 {
					xt = 2
				} else if xt == 1 {
					xt = 3
				}
			} else {
				if xt == 0 {
					xt = 1
				} else if xt == 2 {
					xt = 3
				}
			}
			if (ib.P < ia.P && ib.W <= ia.W) || (ib.P <= ia.P && ib.W < ia.W) {
				if xb == 0 {
					xb = 2
				} else if xb == 1 {
					xb = 3
				}
			} else {
				if xb == 0 {
					xb = 1
				} else if xb == 2 {
					xb = 3
				}
			}
			if xt == xb && xt == 3 {
				break
			}
		}
		s := 0
		if xt == 1 {
			nt++
			s = 1
		}
		if xb == 1 {
			nb++
			if s == 0 {
				s = 2
			} else {
				s = 3
			}
		}
		if Debug && BruteOut != nil {
			fmt.Fprintf(BruteOut, "%d,%d,%d,%d\n", ic, ia.P, ia.W, s)
		}
	}

	return
}
