// Use of this code is governed by the GPL version 2 or later.
// See the file LICENSE for details.

package main

import (
	"log"
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"fmt"
	. "math"
)

var Debug = false

var Binom = make([]int, 10000000)

func computeBinomial() {
	w := []int64{1}
	wn := make([]int64, 0, 10000)

	for n := 1; n < len(Binom); n++ {
		wn = wn[:0]
		t := int64(0)
		for _, e := range w {
			wn = append(wn, e+t)
			if e+t >= int64(len(Binom)) {
				break
			} else {
				if Binom[t+e] == 0 {
					Binom[t+e] = n
				}
			}
			t = e
		}
		if n&1 == 1 {
			wn = append(wn, wn[len(wn)-1])
		}
		w, wn = wn, w
	}
}

func init() {
	log.SetFlags(log.Lshortfile)

	computeBinomial()
}

type Case struct {
	Case int // case #N
	N    int
}

func Parse() []*Case {
	in := bufio.NewReader(os.Stdin)
	buf, err := ioutil.ReadAll(in)
	if err != nil {
		log.Panic(err)
	}
	tok := strings.Fields(string(buf))
	N, err := strconv.Atoi(tok[0])
	if err != nil {
		log.Panic(err)
	}
	if N > 20 || len(tok) != N+1 {
		log.Print("Bad Input N ", N, " len(tok) ", len(tok))
	}

	cases := make([]*Case, 0, N)
	for i := 0; i < N; i++ {
		n, err := strconv.Atoi(tok[i+1])
		if err != nil {
			log.Panic("i ", i, err)
		}
		cases = append(cases, &Case{Case: i + 1, N: n})
	}

	return cases
}

func Steps(s int) int {
	// if we don't find a divisor (eg s is prime) it's
	// s+1 steps

	// I could do a prime check but brute force is pretty fast.
	steps := s + 1
	for d1 := int(Sqrt(float64(s))); d1 > 1; d1-- {
		// log.Print(d1, s/d1, s%d1, Binom[d1], Binom[s/d1])
		if s%d1 == 0 && Binom[d1]+Binom[s/d1] < steps {
			steps = Binom[d1] + Binom[s/d1]
		}
	}

	return steps
}

func main() {
	cases := Parse()
	for _, c := range cases {
		fmt.Println("Case #", c.Case, ": ", Steps(c.N)) //, c.N, Steps(c.N) > c.N)
	}
}
