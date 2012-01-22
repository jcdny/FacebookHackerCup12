// Copyright © 2012 Jeffrey Davis <jeff.davis@gmail.com>
// Use of this code is governed by the GPL version 2 or later.
// See the file LICENSE for details.

package main

func Gcd(n, m int64) int64 {
	for m != 0 {
		n, m = m, n%m
	}
	return n
}

func Lcm(n, m int64) int64 {
	return m / Gcd(n, m) * n
}

func MaxV(v1 int64, vn ...int64) (m int64) {
	m = v1
	for _, vi := range vn {
		if vi > m {
			m = vi
		}
	}
	return
}
