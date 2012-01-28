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
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type Case struct {
	N int // case #N
	M int
	E string
}

const MOD = 4207849484

func Count(m int, s string, memo map[string]int64) int64 {
	//log.Print("m,s: ", m, " ", s)
	if n, ok := memo[s]; ok {
		// log.Print("memo ", n, " \"", s, "\"")
		return n
	}
	so := s
	nc := int64(0)
	c := 0
	for len(s) > 0 {
		n := int(s[0] - '0')
		c = c*10 + n
		if c > m {
			break
		} else {
			s = s[1:]
			if len(s) == 0 {
				nc++
			} else if s[0] != '0' {
				nc += Count(m, s, memo)
			}
		}
	}
	nc = nc % MOD
	memo[so] = nc

	return nc
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
	if N > 25 || len(tok) < 2*N+1 {
		log.Print("Bad Input N ", N, " len(tok) ", len(tok))
	}

	cases := make([]*Case, 0, N)
	for i := 0; i < N; i++ {
		m, err := strconv.Atoi(tok[i*2+1])
		if err != nil {
			log.Panic("i ", i, err)
		}
		cases = append(cases, &Case{N: i + 1, M: m, E: tok[i*2+2]})
	}
	return cases
}

func main() {
	cases := Parse()
	for _, c := range cases {
		memo := make(map[string]int64, len(c.E)+1)
		fmt.Print("Case #", c.N, ": ", Count(c.M, c.E, memo), "\n")
		log.Print(c)
	}
}
