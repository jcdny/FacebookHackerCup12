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
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Cups(s string) int {
	var cc [256]int

	for _, c := range s {
		cc[c]++
	}

	// we need 2 C's so fix this and record this as the current max.
	cc['C'] = cc['C'] / 2
	ncup := cc['C']

	for _, c := range "HAKERUP" {
		if ncup > cc[c] {
			ncup = cc[c]
		}
		if ncup == 0 {
			// any missing letter means no cup for you...
			break
		}
	}

	return ncup
}

func main() {
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}
	line = strings.TrimSpace(line)
	cases, err := strconv.Atoi(line)
	if err != nil {
		log.Panic(err)
	}
	if cases < 1 || cases > 20 {
		log.Panic("cases expected 1 - 20")
	}
	for i := 0; i < cases; i++ {
		line, err := in.ReadString('\n')
		if err != nil {
			log.Panic(err)
		}

		ncup := Cups(line)

		fmt.Fprint(os.Stdout, "Case #", i+1, ": ", ncup, "\n")
	}
}
