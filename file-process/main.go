package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"sync"
)

type result struct {
	NumRows           int
	PeopleCount       int
	CommonName        string
	CommonNameCount   int
	DonationMonthFreq map[string]int
}
type parsed struct {
	fullName string
	date     string
}
type subTotal struct {
	numRows           int
	nameCount         map[string]int
	fullNameCount     map[string]int
	donationMonthFreq []int
}

func mostCommon(nameCount map[string]int) (string, int) {
	var max int
	var maxKey string
	for k, v := range nameCount {
		if v > max {
			max = v
			maxKey = k
		}
	}
	return maxKey, max
}

func main() {

	res := process("./data/test.txt")

	js, err := json.Marshal(res)

	if err != nil {
		fmt.Printf("%s", err.Error())

	} else {
		fmt.Printf("%s", js)
	}
}

func process(path string) result {
	res := result{DonationMonthFreq: map[string]int{}}
	lines := make(chan parsed)
	results := make(chan subTotal)
	var wg sync.WaitGroup

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go worker(lines, results, &wg)
	}
	go read(path, lines)
	go workerWatcher(results, &wg)

	return summarize(results, res)
}

func summarize(results chan subTotal, res result) result {
	finalSubtotal := subTotal{
		numRows:           0,
		nameCount:         map[string]int{},
		fullNameCount:     map[string]int{},
		donationMonthFreq: make([]int, 13)}

	for currentSubtotal := range results {
		finalSubtotal.numRows += currentSubtotal.numRows
		for k, v := range currentSubtotal.nameCount {
			finalSubtotal.nameCount[k] += v
		}
		for k, v := range currentSubtotal.fullNameCount {
			finalSubtotal.fullNameCount[k] += v
		}
		for k, v := range currentSubtotal.donationMonthFreq {
			finalSubtotal.donationMonthFreq[k] += v
		}
	}
	res.NumRows = finalSubtotal.numRows
	res.PeopleCount = len(finalSubtotal.fullNameCount)
	res.CommonName, res.CommonNameCount = mostCommon(finalSubtotal.nameCount)
	res.DonationMonthFreq = convertToMap(finalSubtotal.donationMonthFreq)
	return res
}

func convertToMap(i []int) map[string]int {
	m := make(map[string]int)
	for k, v := range i {
		if k != 0 {
			m[fmt.Sprintf("%02d", k)] = v
		}
	}
	return m
}

func workerWatcher(results chan subTotal, wg *sync.WaitGroup) {
	wg.Wait()
	close(results)
}
