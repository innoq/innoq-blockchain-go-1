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

	actual, expected := generateProofFast(block), uint64(1917336)

	if actual != expected {
		t.Errorf("% q != %q", actual, expected)
	}

}
