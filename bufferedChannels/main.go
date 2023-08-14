package main

import "fmt"

func main() {
	size := 4
	ch := make(chan string, size)

	send(ch, "one")
	send(ch, "two")
	send(ch, "three")
	send(ch, "four")

	fmt.Println("sending done, receiving...")

	for i := 0; i < size; i++ {
		fmt.Println(<-ch)
	}
}

func send(ch chan string, v string) {
	ch <- v
}
