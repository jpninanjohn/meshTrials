package main

import (
	"fmt"
	"log"
	"time"

	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	_, _ = fmt.Fprintf(w, "Hello from service B - V2\n")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
