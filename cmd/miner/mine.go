package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strings"
)

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
