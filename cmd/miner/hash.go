package main

import (
	"crypto/sha256"
	"encoding/json"
)

func hashBlock(block Block) string {
	str, _ := json.Marshal(block)
	sum := sha256.Sum256(str)
	return string(sum[:])
}
