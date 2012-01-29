package main

import (
	"testing"
	"rand"
	"log"
)

func TestChecksum(t *testing.T) {
	if Checksum([]int{2, 1}) != 1024 ||
		Checksum([]int{1, 2}) != 994 ||
		Checksum([]int{2, 3, 5, 1, 4}) != 570316 ||
		Checksum([]int{2, 4, 3, 1}) != 987041 {
		t.Fail()
	}
}

var Dm = false

func mergesort(arr []int, d []byte) ([]int, []byte) {
	n := len(arr)
	if n <= 1 {
		return arr, d
	}
	mid := n / 2
	var first, second []int
	first, d = mergesort(arr[:mid], d)
	second, d = mergesort(arr[mid:], d)

	dl := len(d)
	out, d := merge(first, second, d)
	if Dm {
		log.Print("MS: ", first, second, out, " ", string(d[dl:]))
	}

	return out, d
}

func merge(arr1, arr2 []int, d []byte) ([]int, []byte) {
	if Dm {
		// log.Print("Merge ", arr1, arr2, " was ", string(d))
	}
	r := make([]int, 0, len(arr1)+len(arr2))
	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0] < arr2[0] {
			d = append(d, '1')
			r = append(r, arr1[0])
			arr1 = arr1[1:]
		} else {
			d = append(d, '2')
			r = append(r, arr2[0])
			arr2 = arr2[1:]
		}
	}

	r = append(r, arr1...)
	r = append(r, arr2...)

	return r, d
}

func TestGen(t *testing.T) {
	arr := []int{2, 4, 3, 1}
	_, d := mergesort(arr, []byte{})
	if string(d) != "12212" {
		t.Fail()
	}
}

func xTestGen3(t *testing.T) {
	for np := range Perm5 {
		for i := range Perm5[np] {
			Perm5[np][i]++
		}
	}
	p := make([]int, 5)
	for _, po := range Perm5[1:] {
		copy(p, po)
		_, d := mergesort(p, []byte{})
		ds := d
		_, _, dp := partition(5, d, [][]byte{})
		log.Print(ds, dp)
	}
}

func TestAllPerm5(t *testing.T) {
	for np := range Perm5 {
		for i := range Perm5[np] {
			Perm5[np][i]++
		}
	}
	p := make([]int, 5)
	for np, po := range Perm5[1:] {
		copy(p, po)
		pco := Checksum(p)
		a, d := mergesort(p, []byte{})
		_, _, part := partition(5, d, [][]byte{})
		pn, _ := unsort(a, part)
		pcn := Checksum(pn)
		if pco != pcn {
			t.Log("With ", np, po, " ", string(d))
			t.Log("Got: ", np, pco, pcn, po, pn)
			t.Fail()
		}
	}
}

func TestRandom(t *testing.T) {
	for N := 2; N < 10000; N++ {
		p := rand.Perm(N)
		po := make([]int, N)
		for i := range p {
			p[i]++
		}
		copy(po, p)

		pco := Checksum(p)
		_, d := mergesort(p, []byte{})
		pn := Unsort(N, string(d))
		pcn := Checksum(pn)
		log.Print("Got: ", N, pco, pcn)
		if pco != pcn {
			if len(po) < 100 {
				t.Log("With ", po, " ", string(d))
			}
			t.Fail()
		}
	}
}

var Perm5 = [][]int{
	{0, 1, 2, 3, 4},
	{0, 1, 2, 4, 3},
	{0, 1, 3, 2, 4},
	{0, 1, 3, 4, 2},
	{0, 1, 4, 2, 3},
	{0, 1, 4, 3, 2},
	{0, 2, 1, 3, 4},
	{0, 2, 1, 4, 3},
	{0, 2, 3, 1, 4},
	{0, 2, 3, 4, 1},
	{0, 2, 4, 1, 3},
	{0, 2, 4, 3, 1},
	{0, 3, 1, 2, 4},
	{0, 3, 1, 4, 2},
	{0, 3, 2, 1, 4},
	{0, 3, 2, 4, 1},
	{0, 3, 4, 1, 2},
	{0, 3, 4, 2, 1},
	{0, 4, 1, 2, 3},
	{0, 4, 1, 3, 2},
	{0, 4, 2, 1, 3},
	{0, 4, 2, 3, 1},
	{0, 4, 3, 1, 2},
	{0, 4, 3, 2, 1},
	{1, 0, 2, 3, 4},
	{1, 0, 2, 4, 3},
	{1, 0, 3, 2, 4},
	{1, 0, 3, 4, 2},
	{1, 0, 4, 2, 3},
	{1, 0, 4, 3, 2},
	{1, 2, 0, 3, 4},
	{1, 2, 0, 4, 3},
	{1, 2, 3, 0, 4},
	{1, 2, 3, 4, 0},
	{1, 2, 4, 0, 3},
	{1, 2, 4, 3, 0},
	{1, 3, 0, 2, 4},
	{1, 3, 0, 4, 2},
	{1, 3, 2, 0, 4},
	{1, 3, 2, 4, 0},
	{1, 3, 4, 0, 2},
	{1, 3, 4, 2, 0},
	{1, 4, 0, 2, 3},
	{1, 4, 0, 3, 2},
	{1, 4, 2, 0, 3},
	{1, 4, 2, 3, 0},
	{1, 4, 3, 0, 2},
	{1, 4, 3, 2, 0},
	{2, 0, 1, 3, 4},
	{2, 0, 1, 4, 3},
	{2, 0, 3, 1, 4},
	{2, 0, 3, 4, 1},
	{2, 0, 4, 1, 3},
	{2, 0, 4, 3, 1},
	{2, 1, 0, 3, 4},
	{2, 1, 0, 4, 3},
	{2, 1, 3, 0, 4},
	{2, 1, 3, 4, 0},
	{2, 1, 4, 0, 3},
	{2, 1, 4, 3, 0},
	{2, 3, 0, 1, 4},
	{2, 3, 0, 4, 1},
	{2, 3, 1, 0, 4},
	{2, 3, 1, 4, 0},
	{2, 3, 4, 0, 1},
	{2, 3, 4, 1, 0},
	{2, 4, 0, 1, 3},
	{2, 4, 0, 3, 1},
	{2, 4, 1, 0, 3},
	{2, 4, 1, 3, 0},
	{2, 4, 3, 0, 1},
	{2, 4, 3, 1, 0},
	{3, 0, 1, 2, 4},
	{3, 0, 1, 4, 2},
	{3, 0, 2, 1, 4},
	{3, 0, 2, 4, 1},
	{3, 0, 4, 1, 2},
	{3, 0, 4, 2, 1},
	{3, 1, 0, 2, 4},
	{3, 1, 0, 4, 2},
	{3, 1, 2, 0, 4},
	{3, 1, 2, 4, 0},
	{3, 1, 4, 0, 2},
	{3, 1, 4, 2, 0},
	{3, 2, 0, 1, 4},
	{3, 2, 0, 4, 1},
	{3, 2, 1, 0, 4},
	{3, 2, 1, 4, 0},
	{3, 2, 4, 0, 1},
	{3, 2, 4, 1, 0},
	{3, 4, 0, 1, 2},
	{3, 4, 0, 2, 1},
	{3, 4, 1, 0, 2},
	{3, 4, 1, 2, 0},
	{3, 4, 2, 0, 1},
	{3, 4, 2, 1, 0},
	{4, 0, 1, 2, 3},
	{4, 0, 1, 3, 2},
	{4, 0, 2, 1, 3},
	{4, 0, 2, 3, 1},
	{4, 0, 3, 1, 2},
	{4, 0, 3, 2, 1},
	{4, 1, 0, 2, 3},
	{4, 1, 0, 3, 2},
	{4, 1, 2, 0, 3},
	{4, 1, 2, 3, 0},
	{4, 1, 3, 0, 2},
	{4, 1, 3, 2, 0},
	{4, 2, 0, 1, 3},
	{4, 2, 0, 3, 1},
	{4, 2, 1, 0, 3},
	{4, 2, 1, 3, 0},
	{4, 2, 3, 0, 1},
	{4, 2, 3, 1, 0},
	{4, 3, 0, 1, 2},
	{4, 3, 0, 2, 1},
	{4, 3, 1, 0, 2},
	{4, 3, 1, 2, 0},
	{4, 3, 2, 0, 1},
	{4, 3, 2, 1, 0},
}
