package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", http.StripPrefix("/", fs))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
