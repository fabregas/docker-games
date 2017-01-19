package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("starting server")
	http.ListenAndServe("0.0.0.0:6600", nil)
}
