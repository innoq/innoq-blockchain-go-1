package main

import (
	"log"
	"net/http"
)

func main() {
	chain := NewChain()
	miner := NewMiner(chain, "00")
	overview := NewOverview(chain)

	http.HandleFunc("/", overview.serveJson)

	http.HandleFunc("/mine", miner.mine)

	http.HandleFunc("/blocks", chain.serveJson)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
