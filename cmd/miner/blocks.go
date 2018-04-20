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

	genesis := []byte(GenesisBlock)
	emptyBlock := Block{}
	if err := json.Unmarshal(genesis, &emptyBlock); err != nil {
		panic(err)
	}

	return &Chain{
		blocks: []Block{emptyBlock},
	}
}
