package main

type Transactions struct {
	Id        string `json:"id"`
	Payload   string `json:"payload"`
	Timestamp uint64 `json:"timestamp"`
}
type Block struct {
	Index             uint64         `json:"index"`
	PreviousBlockHash string         `json:"previousBlockHash"`
	Proof             uint64         `json:"proof"`
	Timestamp         uint64         `json:"timestamp"`
	Transactions      []Transactions `json:"transactions"`
}

const GenesisBlock = `{"index":1,"timestamp":0,"proof":1917336,"transactions":[{"id":"b3c973e2-db05-4eb5-9668-3e81c7389a6d","timestamp":0,"payload":"I am Heribert Innoq"}],"previousBlockHash":"0"}`


