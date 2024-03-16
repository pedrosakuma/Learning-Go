package main

import (
	"bufio"
	"bytes"
	"os"
)

func read(file string, output chan parsed) {
	f, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	defer close(output)

	pipe := []byte{byte('|')}

	for scanner.Scan() {
		b := scanner.Bytes()
		// split does not allocate, only slices the original byte slice
		value := bytes.Split(b, pipe)
		// can't avoid allocation here, scanner buffer is reused
		output <- parsed{fullName: string(value[7]), date: string(value[13])}
	}
}
