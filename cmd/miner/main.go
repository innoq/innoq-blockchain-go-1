package main

import (
	"log"
	"net/http"

	sse "github.com/ouven/ssehandler-go"
)

func main() {
	chain := NewChain()
	miner := NewMiner(chain, "00000")
	overview := NewOverview(chain)

	miner.Start()
	defer miner.Stop()

	ssehandler := sse.NewSSEHandler()
	ssehandler.Start()
	defer ssehandler.Stop()

	http.HandleFunc("/", overview.serveJson)

	http.HandleFunc("/mine", miner.mine)

	http.HandleFunc("/blocks", chain.serveJson)

	http.Handle("/events", ssehandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
