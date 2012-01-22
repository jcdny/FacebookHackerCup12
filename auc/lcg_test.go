package main

import (
	"testing"
	"rand"
)

func genLcg(mmax, cmax int64) (m, a, c int64) {
	m = rand.Int63n(mmax-1) + 1
	a = rand.Int63n(cmax)
	c = rand.Int63n(cmax)

	return
}

func BenchmarkLcgOrbit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m, a, c := genLcg(mMax, gMax)
		p := rand.Int63n(m-1) + 1
		//skip, cyc := 
		LcgOrbit(p, m, a, c, 1000000)
		// log.Print("p ", p, " skip ", skip, " len(cyc) ", len(cyc), " cyc [:2] ", cyc[:2])
	}

}
