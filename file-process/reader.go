package main

import (
	"bufio"
	"context"
	"os"
)

var batchSize = 100

func Read(file string, ctx context.Context, rowsBatch *[]string) <-chan []string {
	out := make(chan []string)

	f, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	go func() {
		defer close(out)

		for {
			scanned := scanner.Scan()
			select {
			case <-ctx.Done():
				return

			default:
				row := scanner.Text()

				if len(*rowsBatch) == batchSize || !scanned {
					out <- *rowsBatch
					*rowsBatch = []string{}
				}

				*rowsBatch = append(*rowsBatch, row)
			}

			if !scanned {
				return
			}
		}

	}()

	return out
}
