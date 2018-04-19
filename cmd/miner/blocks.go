package main

import (
	"encoding/json"
	"net/http"
)

type Chain struct {
	blocks []Block
}

func (o *Chain) LastBlock() *Block {
	return &o.blocks[0]
}

func (o *Chain) addBlock(b Block) {
	o.blocks = append([]Block{b}, o.blocks...)
}

func (o *Chain) Height() uint64 {
	return uint64(len(o.blocks))
}

func (o *Chain) Blocks() []Block {
	return o.blocks
}

func (o *Chain) serveJson(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(o.blocks)
}

func NewChain() *Chain {
	genesis := `{
    "index": 1,
    "timestamp": 0,
    "proof": 1917336,
    "transactions": [
        {
            "id": "b3c973e2-db05-4eb5-9668-3e81c7389a6d",
            "timestamp": 0,
            "payload": "I am Heribert Innoq"
        }
    ],
    "previousBlockHash": "0"
	}`

	byt := []byte(genesis)
	dat := Block{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	return &Chain{
		blocks: []Block{dat},
	}
}
