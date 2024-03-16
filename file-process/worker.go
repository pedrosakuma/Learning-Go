package main

import (
	"strconv"
	"strings"
	"sync"
)

func worker(input chan parsed, results chan subTotal, wg *sync.WaitGroup) {

	defer wg.Done()
	subTotal := subTotal{donationMonthFreq: make([]int, 13), fullNameCount: map[string]int{}, nameCount: map[string]int{}}
	for row := range input {
		firstName, fullName, months := processRow(row)
		subTotal.fullNameCount[fullName]++
		subTotal.nameCount[firstName]++
		subTotal.donationMonthFreq[months]++
		subTotal.numRows++
	}
	results <- subTotal
}

func processRow(row parsed) (firstName string, fullName string, month int) {
	// extract full name
	fullName = strings.Replace(strings.TrimSpace(row.fullName), " ", "", -1)

	// extract first name
	name := strings.TrimSpace(row.fullName)
	if name != "" {
		startOfName := strings.Index(name, ", ") + 2
		if endOfName := strings.Index(name[startOfName:], " "); endOfName < 0 {
			firstName = name[startOfName:]
		} else {
			firstName = name[startOfName : startOfName+endOfName]
		}
		if strings.HasSuffix(firstName, ",") {
			firstName = strings.Replace(firstName, ",", "", -1)
		}
	}

	// extract month
	date := strings.TrimSpace(row.date)
	if len(date) == 8 {
		month, _ = strconv.Atoi(date[:2])
	} else {
		month = 0
	}

	return firstName, fullName, month
}
