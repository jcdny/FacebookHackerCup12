package main

import (
	"testing"
	"rand"
	"log"
)

func genCases(n int, cycle int64) []*Auction {
	av := make([]*Auction, 0, 3*n)
	j := 0
	for i := 0; i < n; i++ {
		var m, k, nc int64
		for nc = cycle + 1; nc > cycle; {
			m = (rand.Int63n(int64(cycle/3)) + 1) % mMax
			k = (rand.Int63n(int64(cycle/3)) + 1) % mMax
			nc = Lcm(m, k) - 1
		}
		a := rand.Int63n(gMax)
		b := rand.Int63n(gMax)
		c := rand.Int63n(gMax)
		d := rand.Int63n(gMax)
		p1 := rand.Int63n(m-1) + 1
		w1 := rand.Int63n(k-1) + 1
		av = append(av, &Auction{Case: j, N: int64(nc/10) + 2, P1: p1, W1: w1, M: m, K: k, A: a, B: b, C: c, D: d})
		j++
		av = append(av, &Auction{Case: j, N: int64(nc * 2), P1: p1, W1: w1, M: m, K: k, A: a, B: b, C: c, D: d})
		j++
		av = append(av, &Auction{Case: j, N: int64(nc + nc/2), P1: p1, W1: w1, M: m, K: k, A: a, B: b, C: c, D: d})
		j++
		// this won't work for brute force check obviously - need to verify no overflows though
		//av = append(av, &Auction{Case: j, N: int64(1e18), P1: p1, W1: w1, M: m, K: k, A: a, B: b, C: c, D: d})
		//j++
	}

	return av
}

func TestIt(t *testing.T) {
	//mcycle := int64(1e14)
	//av := genCases(3, mcycle)
	mcycle := int64(10000)
	av := genCases(10, mcycle)
	for _, a := range av {
		nt, nb := ComputeCase(a)
		ntb, nbb := nt*0, nt*0
		if mcycle < 100000 {
			ntb, nbb = ComputeCaseBrute(a)
			match := ntb == nt && nbb == nb
			if !match {
				log.Print("------------------------------------------------------------")
				log.Print(a.Case, match, ntb, nbb, nt, nb)
				log.Printf("%#v", a)
				log.Print(a)
			}
			if !match {
				t.Fail()
			}
		}
	}
}

func BenchmarkIt(b *testing.B) {
	mcycle := int64(1e11)
	av := genCases(10, mcycle)
	for i := 0; i < b.N; i++ {
		for _, a := range av {
			ComputeCase(a)
		}
	}
}
