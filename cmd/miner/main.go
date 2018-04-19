package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		b := Block{Index: "1", Timestamp: "0", Proof: 0, Transaction: []string{}, PreviousBlockHash: "0"}
		json.NewEncoder(w).Encode(b)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hash(block Block) {
	str, _ := json.Marshall(block)
}
