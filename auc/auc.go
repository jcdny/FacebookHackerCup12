// Copyright © 2012 Jeffrey Davis <jeff.davis@gmail.com>
// Use of this code is governed by the GPL version 2 or later.
// See the file LICENSE for details.

package main

import (
	"log"
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
	"math"
)

const gMax = int64(1e9)
const nMax = int64(1e18)
const mMax = int64(1e7)

var Debug = false
var BruteOut *os.File

func init() {
	var err os.Error
	log.SetFlags(log.Lshortfile)
	if Debug {
		BruteOut, err = os.Create("brute.csv")
		if err != nil {
			log.Print("Failed to create brute.csv: ", err)
		}
	}
}

type Auction struct {
	Case       int
	N          int64
	P1         int64
	W1         int64
	M, K       int64
	A, B, C, D int64
}

type Item struct {
	P, W int
	N    int
}

type RangeElem struct {
	lo  int64
	nlo int64
	hi  int64
	nhi int64
}
type Range []RangeElem

// Given an i, j and m (i is P_i or W_i depending on which has the
// the smallest range, j is the other) update the range and count at
// the extreme.
func (r Range) Update(i, j, m int64) {
	//	log.Print("UU ", i, j, m)
	//	log.Print(i, j, len(r))
	if r[i].lo > j || r[i].lo == 0 {
		r[i].lo = j
		r[i].nlo = m
	} else if r[i].lo == j {
		r[i].nlo += m
	}
	if r[i].hi > j {
	} else if r[i].hi < j {
		r[i].hi = j
		r[i].nhi = m
	} else {
		r[i].nhi += m
	}
}

func (r Range) Log() {
	t := make(map[int]RangeElem, len(r)/10)
	for i, ri := range r {
		if ri.lo != 0 {
			t[i] = ri
		}
	}

	log.Print(t)
}

func (r Range) Count() (a int64, b int64) {
	//r.Log()
	m := int64(math.MaxInt64)
	for i := 0; i < len(r); i++ {
		if r[i].lo < m && r[i].lo != 0 {
			m = r[i].lo
			a += r[i].nlo
		}
	}

	m = 0
	for i := len(r) - 1; i >= 0; i-- {
		if r[i].hi > m {
			m = r[i].hi
			b += r[i].nhi
		}
	}
	return a, b
}

func LcgOrbit(p, m, a, c, n int64) (skip []int64, cycle []int64, max int64) {
	cv := make([]int64, mMax+1)
	pv := make([]int64, 0, mMax+1)
	max = p

	n += 1
	var j int64
	for j = 1; cv[p] < 1 && j < n; j++ {
		cv[p] = j
		pv = append(pv, p)
		if p > max {
			max = p
		}
		p = ((p*a + c) % m) + 1
	}
	nskip := cv[p] - 1
	ncycle := j - cv[p]

	// forcing generating more here insures the skip lengths are the
	// same.  I am pretty sure the max skip length is ~ 24 .since we
	// peel at least 1 bit per iteration (and that would be wildly
	// improbable unless they construct the rng perversely) but better
	// safe than sorry for now.
	for ; nskip < 64; nskip++ {
		pv = append(pv, p)
		if p > max {
			max = p
		}
		p = ((p*a + c) % m) + 1
	}

	skip = append(skip, pv[:nskip]...)
	cycle = pv[nskip : nskip+ncycle]

	return
}

func ComputeCase(a *Auction) (nt, nb int64) {
	n := a.N
	ps, po, mp := LcgOrbit(a.P1, a.M, a.A, a.B, a.N)
	ws, wo, mw := LcgOrbit(a.W1, a.K, a.C, a.D, a.N)
	r := make(Range, MaxV(mp, mw)+1)
	if len(ps) != len(ws) {
		log.Panic("Skip mismatch")
	}

	// run through skip
	for i := range ps {
		if int64(i) >= n {
			break
		}
		if mw > mp {
			r.Update(ps[i], ws[i], 1) // range on P_i
		} else {
			r.Update(ws[i], ps[i], 1) // range on W_i
		}
	}
	n -= int64(len(ps))

	if n > 0 {
		// now handle cyclic part
		cycle := Lcm(int64(len(po)), int64(len(wo)))
		log.Print(cycle, a)
		mult := n / cycle
		rem := n - mult*cycle
		if rem > 0 {
			mult++
		}
		if Debug {
			log.Print(a.Case, " n ", n, " rem ", rem, " mult ", mult, " cycle ", cycle, " lpo, lwo ", len(po), len(wo))
		}
		rem--
		for i, ip, iw := int64(0), 0, 0; i < cycle; i++ {
			if mw > mp {
				r.Update(po[ip], wo[iw], mult) // range on P_i
			} else {
				r.Update(wo[iw], po[ip], mult) // range on W_i
			}
			if i == rem {
				mult--
				if mult == 0 {
					break
				}
			}
			ip++
			if ip == len(po) {
				ip = 0
			}
			iw++
			if iw == len(wo) {
				iw = 0
			}

		}
	}

	nb, nt = r.Count()

	return
}

func ParseCase(i int, line string) *Auction {
	a := &Auction{Case: i}
	tokens := strings.Fields(line)
	for p, s := range tokens {
		val, err := strconv.Atoi64(s)
		if err != nil {
			log.Panic(i, line, err)
		}
		switch p {
		case 0:
			a.N = val
		case 1:
			a.P1 = val
		case 2:
			a.W1 = val
		case 3:
			a.M = val
		case 4:
			a.K = val
		case 5:
			a.A = val
		case 6:
			a.B = val
		case 7:
			a.C = val
		case 8:
			a.D = val
		}
	}
	if Debug {
		log.Printf("%#v", *a)
	}
	return a
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var line string
	var err os.Error
	for {
		line, err = in.ReadString('\n')
		if err != nil {
			log.Panic(err)
		}
		line = strings.TrimSpace(line)
		if line == "" || line[0] != '#' {
			break
		}
	}

	cases, err := strconv.Atoi(line)
	if err != nil {
		log.Panic(err)
	}
	if cases < 1 || cases > 20 {
		log.Panic("cases expected 1 - 20")
	}

	if Debug && BruteOut != nil {
		// header line for the brute force csv
		fmt.Fprint(BruteOut, "case,p,w,s\n")
	}

	for i := 0; i < cases; i++ {
		line, err := in.ReadString('\n')
		if err != nil {
			log.Panic(err)
		}
		a := ParseCase(i+1, line)
		if Debug {
			ntb, nbb := ComputeCaseBrute(a)
			fmt.Fprint(os.Stderr, "Brute Case #", a.Case, ": ", ntb, nbb, "\n")
		}
		nt, nb := ComputeCase(a)

		log.Print("Case #", a.Case, ": ", nt, nb, "\n")
		fmt.Fprint(os.Stdout, "Case #", a.Case, ": ", nt, nb, "\n")
	}
}
