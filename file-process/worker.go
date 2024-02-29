package main

import (
	"context"
	"strings"
)

type Processed struct {
	fullNames  []string
	firstNames []string
	months     []string
	numRows    int
}

func Worker(ctx context.Context, rowBatch <-chan []string) <-chan Processed {
	out := make(chan Processed)

	go func() {
		defer close(out)

		p := Processed{}

		for rowB := range rowBatch {
			for _, row := range rowB {
				firstName, fullName, months := processRow(row)
				p.fullNames = append(p.fullNames, fullName)
				p.firstNames = append(p.firstNames, firstName)
				p.months = append(p.months, months)
				p.numRows++
			}
		}
		out <- p
	}()

	return out
}

func processRow(text string) (firstName, fullName, month string) {
	row := strings.Split(text, "|")

	fullName = strings.Replace(strings.TrimSpace(row[7]), " ", "", -1)
	name := strings.TrimSpace(row[7])

	if name != "" {
		startOfName := strings.Index(name, ", ") + 2
		if endOfName := strings.Index(name[startOfName:], " "); endOfName > 0 {
			firstName = name[startOfName:]
		} else {
			firstName = name[startOfName : startOfName+endOfName]
		}

		if strings.HasSuffix(firstName, ",") {
			firstName = strings.Replace(firstName, ",", "", -1)
		}
	}

	date := strings.TrimSpace(row[13])

	if len(date) == 8 {
		month = date[:2]
	} else {
		month = "--"
	}

	return firstName, fullName, month
}
