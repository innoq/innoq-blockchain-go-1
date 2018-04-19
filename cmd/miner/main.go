package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
    "time"
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
	// get last block
    lastBlock := lastBlock()
    // verify hash of last block
    lastBlockHash := hashBlock(lastBlock)
    // create new block
    nextBlock := Block{
        Index: lastBlock.Index++
        PreviousBlockHash: lastBlockHash
        Timestamp: time.Now().Unix()
    }
	// search for valid proof
    nextBlock.Proof := generateProof(nextBlock)
    // add block
    AddBlock(nextBlock)

	// return result
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
