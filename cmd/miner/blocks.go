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
	w.Header().Set("Content-Type", "application/json")
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

func findTransactionById(o *Chain, transactionId string) *Transaction {
	// empty Chain -> return nil
	if len(o.blocks) == 0 {
		return nil
	}

	// iterate over blocks
	for _, currentBlock := range o.blocks {

		// iterate over transactions in payload
		for _, currentTransaction := range currentBlock.Transactions {
			if currentTransaction.Id == transactionId {
				return &currentTransaction
			}
		}

	}

	// seems we didn't find anything...
	return nil
}
