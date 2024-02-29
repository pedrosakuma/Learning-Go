package main

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
)

type result struct {
	numRows           int
	peopleCount       int
	commonName        string
	commonNameCount   int
	donationMonthFreq map[string]int
}

func main() {
	res := result{donationMonthFreq: map[string]int{}}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	rowsBatch := []string{}
	rowsCh := Read("./data/test.txt", ctx, &rowsBatch)

	workersCh := make([]<-chan Processed, runtime.NumCPU())

	for i := 0; i < runtime.NumCPU(); i++ {
		workersCh[i] = Worker(ctx, rowsCh)
	}

	firstNameCount := map[string]int{}
	fullNameCount := map[string]bool{}

	for processed := range Combiner(ctx, workersCh...) {
		// add number of rows processed by worker
		res.numRows += processed.numRows

		// add months processed by worker
		for _, month := range processed.months {
			res.donationMonthFreq[month]++
		}

		// use full names to count people
		for _, fullName := range processed.fullNames {
			fullNameCount[fullName] = true
		}
		res.peopleCount = len(fullNameCount)

		// update most common first name based on processed results
		for _, firstName := range processed.firstNames {
			firstNameCount[firstName]++

			if firstNameCount[firstName] > res.commonNameCount {
				res.commonName = firstName
				res.commonNameCount = firstNameCount[firstName]
			}
		}
	}

	js, err := json.Marshal(res)

	if err != nil {
		fmt.Printf("%s", err.Error())

	} else {
		fmt.Printf("%s", js)
	}
}
