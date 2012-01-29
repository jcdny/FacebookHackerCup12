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

var Du = false // debug

func init() {
	log.SetFlags(log.Lshortfile)
}

type Case struct {
	Case int // case #N
	N    int
	D    string
}

func Checksum(arr []int) int {
	r := 1
	for _, a := range arr {
		r = (31*r + a) % 1000003
	}
	return r
}

func Ordered(n int) []int {
	sa := make([]int, n)
	for i := range sa {
		sa[i] = i + 1
	}
	return sa
}

func Unsort(n int, d string) []int {
	_, _, part := partition(n, []byte(d), [][]byte{})
	arr := Ordered(n)
	out, _ := unsort(arr, part)
	return out
}

// partition takes N and D and generates the partitioned version of D.	
func partition(arr int, d []byte, dp [][]byte) (int, []byte, [][]byte) {
	if arr <= 1 {
		return arr, d, dp
	}
	mid := arr / 2
	var first, second int
	first, d, dp = partition(mid, d, dp)
	second, d, dp = partition(arr-mid, d, dp)
	out, d, do := partitionit(first, second, d)
	dp = append(dp, do)

	return out, d, dp
}

func partitionit(arr1, arr2 int, d []byte) (int, []byte, []byte) {
	do := []byte{}
	r := arr1 + arr2
	for arr1 > 0 && arr2 > 0 {
		if d[0] == '1' {
			arr1--
		} else {
			arr2--
		}
		do = append(do, d[0])
		d = d[1:]
	}

	if Du {
		log.Print("Partition \"", string(do), "\" leaves ", string(d))
	}
	return r, d, do
}

func unsort(arr []int, part [][]byte) ([]int, [][]byte) {
	if len(arr) <= 1 {
		return arr, part
	}

	first, second, part := unmerge(arr, part)
	second, part = unsort(second, part)
	first, part = unsort(first, part)

	out := append(first, second...)

	return out, part
}

func unmerge(arr []int, part [][]byte) ([]int, []int, [][]byte) {
	n := len(arr)
	l1 := n / 2
	l2 := n - l1

	dmove := part[len(part)-1]
	part = part[:len(part)-1]

	var arr1, arr2 []int
	for i, dc := range dmove {
		if dc == '1' {
			l1--
			arr1 = append(arr1, arr[i])
		} else {
			l2--
			arr2 = append(arr2, arr[i])
		}
	}
	if l1 > 0 {
		arr1 = append(arr1, arr[len(dmove):]...)
	}
	if l2 > 0 {
		arr2 = append(arr2, arr[len(dmove):]...)
	}

	return arr1, arr2, part
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
	if N > 20 || len(tok) != 2*N+1 {
		log.Print("Bad Input N ", N, " len(tok) ", len(tok))
	}

	cases := make([]*Case, 0, N)
	for i := 0; i < N; i++ {
		n, err := strconv.Atoi(tok[i*2+1])
		if err != nil {
			log.Panic("i ", i, err)
		}
		cases = append(cases, &Case{Case: i + 1, N: n, D: tok[i*2+2]})
	}
	return cases
}

func main() {
	cases := Parse()
	for _, c := range cases {
		arr := Unsort(c.N, c.D)
		chk := Checksum(arr)
		fmt.Print("Case #", c.Case, ": ", chk, "\n")
	}
}
