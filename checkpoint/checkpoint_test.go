package main

import (
	"testing"
	// "rand"
	"log"
)

func TestSteps(t *testing.T) {
	for i := 10000000; i > 0; i-- {
		log.Print(i, Steps(i))
	}
}
