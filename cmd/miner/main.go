package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func main() {
	chain := NewChain()
	overview := NewOverview(chain)

	http.HandleFunc("/", overview.serveJson)

	http.HandleFunc("/mine", mine)

	http.HandleFunc("/blocks", blocks)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mine(w http.ResponseWriter, r *http.Request) {
	// get last block
	lastBlock := Block{} //lastBlock()
	// verify hash of last block
	lastBlockHash := hashBlock(lastBlock)

	// create new block
	nextBlock := Block{
		Index:             lastBlock.Index + 1,
		PreviousBlockHash: lastBlockHash,
		Timestamp:         time.Now().Unix(),
	}

	// search for valid proof
	nextBlock.Proof = generateProof(nextBlock, "00")
	// add block
	// AddBlock(nextBlock)

	// return result
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
