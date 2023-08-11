package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	ch := make(chan string)

	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	for _, api := range apis {
		go checkApi(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Printf(<-ch)
	}

	fmt.Printf("Done It took %v milliseconds /n", time.Since(start).Milliseconds())

}

func checkApi(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("Error %s is down \n", api)
		return
	}

	ch <- fmt.Sprintf("Success %s is up and running \n", api)
}
