package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	events := NewEvents()
	events.Start()
	defer events.Stop()

	chain := NewChain()
	miner := NewMiner(chain, events, "00000")
	overview := NewOverview(chain)
	transactions := NewTransactions()

	miner.Start()
	defer miner.Stop()

	r.HandleFunc("/", overview.serveJson)

	r.HandleFunc("/mine", miner.mine)

	r.HandleFunc("/blocks", chain.serveJson)

	r.Handle("/events", events)

	r.HandleFunc("/transactions", transactions.Post)

	r.HandleFunc("/transactions/{id}", transactions.serveJson)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
