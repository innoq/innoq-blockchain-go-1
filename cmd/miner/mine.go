package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Miner struct {
	chain  *Chain
	prefix string
}

func NewMiner(chain *Chain, prefix string) *Miner {
	return &Miner{
		chain:  chain,
		prefix: prefix,
	}
}

func generateProof(block Block, prefix string) uint64 {

	var hash string

	for n := uint64(0); !strings.HasPrefix(hash, prefix); n++ {
		block.Proof = n
		str, _ := json.Marshal(block)
		sum := sha256.Sum256([]byte(string(str)))
		hash = hex.EncodeToString(sum[:])
	}

	return block.Proof
}

func (m *Miner) mine(w http.ResponseWriter, r *http.Request) {
	// get last block
	lastBlock := *m.chain.LastBlock()
	// verify hash of last block
	lastBlockHash := hashBlock(lastBlock)

	// create new block
	nextBlock := Block{
		Index:             lastBlock.Index + 1,
		PreviousBlockHash: lastBlockHash,
		Timestamp:         time.Now().Unix(),
		Transactions:      []Transaction{},
	}

	// search for valid proof
	nextBlock.Proof = generateProof(nextBlock, m.prefix)
	// add block
	m.chain.addBlock(nextBlock)

	// return result
	fmt.Fprintf(w, "Hello, miner")
}
