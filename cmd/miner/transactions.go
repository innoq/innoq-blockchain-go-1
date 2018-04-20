package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	uuid "github.com/satori/go.uuid"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Payload struct {
	Payload string `json:"payload"`
}

type Transactions struct {
	pool   []Transaction
	events Events
}

func (transactions *Transactions) Add(transaction Transaction) {
	transactions.pool = append([]Transaction{transaction}, transactions.pool...)
	transactions.events.SendNewTransactionEvent(&transaction)
}

func (transactions *Transactions) Pop() []Transaction {
	return transactions.pool[:min(4, len(transactions.pool))]
}

func (transactions *Transactions) Get(id string) *Transaction {
	for _, t := range transactions.pool {
		if t.Id == id {
			return &t
		}
	}
	return nil
}

func (transactions *Transactions) Create(payload io.Reader) (*Transaction, error) {

	emptyPayload := Payload{}
	if err := json.NewDecoder(payload).Decode(&emptyPayload); err != nil {
		return nil, err
	}

	return &Transaction{
		Id:        uuid.NewV4().String(),
		Timestamp: uint64(time.Now().Unix()),
		Payload:   emptyPayload.Payload,
	}, nil
}

func (transactions *Transactions) Post(w http.ResponseWriter, r *http.Request) {

	wireContext, _ := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))

	span := opentracing.StartSpan("Transactions:Post", ext.RPCServerOption(wireContext))
	defer span.Finish()

	if r.Method == "POST" {
		if transaction, err := transactions.Create(r.Body); err != nil {
			http.Error(w, err.Error(), 400)
		} else {
			transactions.Add(*transaction)
			json.NewEncoder(w).Encode(transaction)
		}
	} else {
		json.NewEncoder(w).Encode(transactions.pool)
	}
}

func (transactions *Transactions) serveJson(w http.ResponseWriter, r *http.Request) {

	wireContext, _ := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))

	span := opentracing.StartSpan("Transactions:serveJson", ext.RPCServerOption(wireContext))
	defer span.Finish()

	vars := mux.Vars(r)
	json.NewEncoder(w).Encode(transactions.Get(vars["id"]))
}

func NewTransactions(events *Events) *Transactions {
	return &Transactions{
		pool:   []Transaction{},
		events: *events,
	}
}
