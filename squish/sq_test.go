package main

import (
	"testing"
	"rand"
	"strconv"
	"log"
)

func genCases(n, lo, hi, l int, valid bool) []*Case {
	cv := make([]*Case, 0, hi-lo)
	for m := lo; m <= hi; m++ {
		for i := 0; i < n; i++ {
			s := ""
			so := ""
			for len(s) < l {
				i := rand.Intn(m) + 1
				if i > m {
					log.Panic("yikes")
				}
				s += strconv.Itoa(i)
				so += strconv.Itoa(i)
			}
			cv = append(cv, &Case{N: m, M: m, E: s})
			cv = append(cv, &Case{N: m, M: m, E: s + "900"})
			cv = append(cv, &Case{N: m, M: m, E: "900" + s})
		}
	}
	return cv

}

func TestIt(t *testing.T) {
	for cc := 990; cc < 991; cc += 7 {
		cv := genCases(5, 2, 255, cc, true)
		log.Print("running ", cc, len(cv))
		for _, c := range cv {
			memo := make(map[string]int64, 1000)
			nc := Count(c.M, c.E, memo)
			log.Print(c.M, nc)
		}
	}
}
