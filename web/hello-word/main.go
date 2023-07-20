package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s", r.URL.Path)
    })

	var porta = ":4474"  	
	http.ListenAndServe(porta, nil)
}