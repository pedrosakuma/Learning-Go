package main

import "testing"

func TestGetVal(t *testing.T) {
	for i := 0; i < 1000; i++ {
		process("./data/test.txt")
	}
}
