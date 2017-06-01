package main

import (
	"log"
	"net/http"
)

func main() {
	r := NewRouter()

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
