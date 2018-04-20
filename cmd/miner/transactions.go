package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type Transactions struct {
	pool []Transaction
}

func (transactions *Transactions) Add(transaction Transaction) {
	transactions.pool = append([]Transaction{transaction}, transactions.pool...)
}

func (transactions *Transactions) Get(id string) *Transaction {
	for _, t := range transactions.pool {
		if t.Id == id {
			return &t
		}
	}
	return nil
}

func (transactions *Transactions) Create(payload string) *Transaction {
	return &Transaction{
		Id:        uuid.NewV4().String(),
		Timestamp: uint64(time.Now().Unix()),
		Payload:   payload,
		Confirmed: false,
	}
}

func (transactions *Transactions) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		transaction := transactions.Create(string(body))
		transactions.Add(*transaction)
		json.NewEncoder(w).Encode(transaction)
	} else {
		json.NewEncoder(w).Encode(transactions.pool)
	}
}

func (transactions *Transactions) serveJson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	json.NewEncoder(w).Encode(transactions.Get(vars["id"]))
}

func NewTransactions() *Transactions {
	return &Transactions{
		pool: []Transaction{},
	}
}
