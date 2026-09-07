package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "{ \"status\" : \"ok\" }")
}

func main() {
	http.HandleFunc("/health", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
