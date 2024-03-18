package main

import (
	"context"
	"testing"
)

func TestGetVal(t *testing.T) {
	ctx := context.Background()

	for i := 0; i < 1000; i++ {
		process(ctx, "./data/test.txt")
	}
}
