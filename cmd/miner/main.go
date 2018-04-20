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

	transactions := NewTransactions(events)

	chain := NewChain()
	miner := NewMiner(chain, events, transactions, "00000")

	overview := NewOverview(chain)
	ui := NewUi(overview)
	transactions := NewTransactions(*events)

	miner.Start()
	defer miner.Stop()

	r.HandleFunc("/", overview.serveJson).Methods(http.MethodGet)

	r.HandleFunc("/mine", miner.mine).Methods(http.MethodGet)

	r.HandleFunc("/blocks", chain.serveJson).Methods(http.MethodGet)

	r.Handle("/events", events).Methods(http.MethodGet)

	r.HandleFunc("/transactions", transactions.Post).Methods(http.MethodGet, http.MethodPost)

	r.HandleFunc("/transactions/{id}", transactions.serveJson).Methods(http.MethodGet)

	r.HandleFunc("/ui", ui.GetIndex).Methods(http.MethodGet)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
