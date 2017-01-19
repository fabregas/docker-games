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
	http.ListenAndServe(":6000", nil)
}
