package main

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
)

type result struct {
	NumRows           int
	PeopleCount       int
	CommonName        string
	CommonNameCount   int
	DonationMonthFreq map[string]int
}

func main() {
	res := result{DonationMonthFreq: map[string]int{}}

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

	// entender melhor os resultados, só copiei e colei
	for processed := range Combiner(ctx, workersCh...) {
		// add number of rows processed by worker
		res.NumRows += processed.numRows

		// add months processed by worker
		for _, month := range processed.months {
			res.DonationMonthFreq[month]++
		}

		// use full names to count people
		for _, fullName := range processed.fullNames {
			fullNameCount[fullName] = true
		}
		res.PeopleCount = len(fullNameCount)

		// update most common first name based on processed results
		for _, firstName := range processed.firstNames {
			firstNameCount[firstName]++

			if firstNameCount[firstName] > res.CommonNameCount {
				res.CommonName = firstName
				res.CommonNameCount = firstNameCount[firstName]
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
