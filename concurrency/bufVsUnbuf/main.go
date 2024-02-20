package main

import "fmt"

func main() {

	ch := make(chan string, 8)

	go ret(ch)
	go ret(ch)
	go ret(ch)
	go ret(ch)

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	go opt(ch)
	go opt(ch)
	go opt(ch)
	go opt(ch)
	go opt(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println()

	// ret(ch)

	// fmt.Println(<-ch)

}

func ret(ch chan string) {
	ch <- "hello"
}

func opt(ch chan string) {
	<-ch
}
