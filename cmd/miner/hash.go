package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

func hashBlock(block Block) string {
	str, _ := json.Marshal(block)
	sum := sha256.Sum256(str)
	return fmt.Sprintf("%x", sum)
}
