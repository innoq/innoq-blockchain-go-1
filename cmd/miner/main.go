package main

import (
	"log"
	"net/http"
)

func main() {
	chain := NewChain()
	miner := NewMiner(chain, "00000")
	overview := NewOverview(chain)

	miner.Start()
	defer miner.Stop()

	http.HandleFunc("/", overview.serveJson)

	http.HandleFunc("/mine", miner.mine)

	http.HandleFunc("/blocks", chain.serveJson)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
