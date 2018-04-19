package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	overview := NewOverview()

	http.HandleFunc("/", overview.serveJson)
	http.HandleFunc("/mine", mine)
	http.HandleFunc("/blocks", blocks)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mine(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
