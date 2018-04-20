package main

type Transaction struct {
	Id        string `json:"id"`
	Timestamp uint64 `json:"timestamp"`
	Payload   string `json:"payload"`
}

type Block struct {
	Index             uint64        `json:"index"`
	Timestamp         int64         `json:"timestamp"`
	Proof             uint64        `json:"proof"`
	Transactions      []Transaction `json:"transactions"`
	PreviousBlockHash string        `json:"previousBlockHash"`
}

const magicGenesisTransactionID = "b3c973e2-db05-4eb5-9668-3e81c7389a6d"

const GenesisBlock = `{
	"index":1,
	"timestamp":0,
	"proof":1917336,
	"transactions":[
		{
			"id":"` + magicGenesisTransactionID +
	`","timestamp":0,
			"payload":"I am Heribert Innoq"
		}
	],
	"previousBlockHash":"0"
	}`
