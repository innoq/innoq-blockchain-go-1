package main

import (
	"net/http"
	"strconv"

	sse "github.com/ouven/ssehandler-go"
)

type Events struct {
	send    chan *sse.Event
	handler sse.SSEHandler
}

func NewEvents() *Events {
	return &Events{
		send:    make(chan *sse.Event, 1),
		handler: sse.NewSSEHandler(),
	}
}

func (e *Events) Start() {
	e.handler.Start()
	go func() {
		lastId := 0
		for {
			event := <-e.send
			lastId = lastId + 1
			event.Id = strconv.Itoa(lastId)
			e.handler.PublishChannel() <- event
		}
	}()
}

func (e *Events) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.handler.ServeHTTP(w, r)
}

func (e *Events) Stop() {
	e.handler.Stop()
}

func (e *Events) SendNewBlockEvent(block *Block) {
	e.send <- &sse.Event{
		Event: "new_block",
		Data:  block,
	}
}

func (e *Events) SendNewTransactionEvent(transaction *Transaction) {
	e.send <- &sse.Event{
		Event: "new_transaction",
		Data:  transaction,
	}
}

func (e *Events) SendNewNodeEvent(node interface{}) {
	e.send <- &sse.Event{
		Event: "new_block",
		Data:  node,
	}
}
