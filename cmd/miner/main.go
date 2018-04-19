package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

type Overview struct {
	NodeId             string `json:"nodeId"`
	CurrentBlockHeight int64  `json:"currentBlockHeight"`
}

func main() {
	http.HandleFunc("/", overview)

	http.HandleFunc("/mine", mine)

	http.HandleFunc("/blocks", blocks)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func overview(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func mine(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func blocks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
