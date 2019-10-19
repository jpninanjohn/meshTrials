package main

import (
	"fmt"
	"log"

	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello from service B - V2\n")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
