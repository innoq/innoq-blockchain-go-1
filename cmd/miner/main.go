package main

import (
	"log"
	"net/http"
)

func main() {
	events := NewEvents()
	events.Start()
	defer events.Stop()

	chain := NewChain()
	miner := NewMiner(chain, events, "00000")
	overview := NewOverview(chain)

	miner.Start()
	defer miner.Stop()

	http.HandleFunc("/", overview.serveJson)

	http.HandleFunc("/mine", miner.mine)

	http.HandleFunc("/blocks", chain.serveJson)

	http.Handle("/events", events)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
