package main

import "fmt"

func main() {
	size := 2
	ch := make(chan string, size) //transform into a buffer, acting like a queue

	send(ch, "one")
	send(ch, "two")
	go send(ch, "three") //iria bloquear
	go send(ch, "four")  //iria bloquear

	fmt.Println(" sending done, receiving...")

	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}
}

func send(ch chan string, v string) {
	ch <- v
}
