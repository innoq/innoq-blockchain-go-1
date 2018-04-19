package main

import (
	"encoding/json"
	"testing"
)

func TestIrgendwas(t *testing.T) {
	byt := []byte(`{"timestamp":0,"index":1,"proof":0,"transactions":[{"id":"b3c973e2-db05-4eb5-9668-3e81c7389a6d","timestamp":0,"payload":"I am Heribert Innoq"}],"previousBlockHash":"0"}`)
	block := Block{}
	if err := json.Unmarshal(byt, &block); err != nil {
		t.Fatal("Unmarshal error")
	}
	proof := generateProof(block, "000000")
	t.Log(proof)
	if proof != 1917336 {
		t.Errorf("fail")
	}
}
