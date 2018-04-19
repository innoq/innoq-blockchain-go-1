package main

import (
	"encoding/json"
	"testing"
)

func TestGenesisBlockYields1917336(t *testing.T) {

	block := Block{}
	if err := json.Unmarshal([]byte(GenesisBlock), &block); err != nil {
		t.Fatal("Unmarshal error")
	}
	proof := generateProof(block, "000000")
	t.Log(proof)
	if proof != 1917336 {
		t.Errorf("fail")
	}
}
