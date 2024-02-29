package main

import (
	"context"
	"sync"
)

func Combiner(ctx context.Context, inputs ...<-chan Processed) <-chan Processed {

	out := make(chan Processed)

	var wg sync.WaitGroup

	multiplexer := func(p <-chan Processed) {
		defer wg.Done()

		for in := range p {
			select {
			case <-ctx.Done():
			case out <- in:
			}
		}
	}

	wg.Add(len(inputs))

	for _, in := range inputs {
		go multiplexer(in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out

}
