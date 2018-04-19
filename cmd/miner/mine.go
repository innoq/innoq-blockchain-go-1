package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type Mine struct {
	answer chan Mined
}
type Mined struct {
	Message string `json:"message"`
	Block   Block  `json:"block"`
}

type Miner struct {
	chain  *Chain
	prefix string
	Queue  chan *Mine
}

func NewMiner(chain *Chain, prefix string) *Miner {
	return &Miner{
		chain:  chain,
		prefix: prefix,
		Queue:  make(chan *Mine, 20),
	}
}

func (m *Miner) Start() {
	go func() {
		for {
			m.findBlock(<-m.Queue)
		}
	}()
}

func (m *Miner) Stop() {}

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

func generateProofFast(block Block) uint64 {

	var sum [32]byte
	block.Proof = ^uint64(0)
	str, _ := json.Marshal(block)
	sum = sha256.Sum256([]byte(string(str)))

	for n := uint64(0); sum[0] != 0 || sum[1] != 0 || sum[2] != 0; n++ {
		block.Proof = n
		str, _ = json.Marshal(block)
		sum = sha256.Sum256([]byte(string(str)))
	}

	return block.Proof
}

func (m *Miner) mine(w http.ResponseWriter, r *http.Request) {
	job := Mine{
		answer: make(chan Mined, 1),
	}
	m.Queue <- &job
	a := <-job.answer
	json.NewEncoder(w).Encode(a)
}

func (m *Miner) findBlock(mine *Mine) {
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
	mine.answer <- Mined{
		Message: "hello miner",
		Block:   nextBlock,
	}
}
