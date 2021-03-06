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

	actual, expected := generateProof(block, "000000"), uint64(1917336)

	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}

	actual, expected = generateProofFast(block, 3), uint64(1917336)

	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}

}

func TestMoreBlocks(t *testing.T) {

	block := Block{
		Timestamp:         42,
		PreviousBlockHash: "01010101",
	}

	for i := 0; i < 10; i++ {
		block.Index = uint64(i)

		block2 := block

		actual := generateProof(block, "0000")
		expected := generateProofFast(block2, 2)

		if actual != expected {
			t.Errorf("%d != %d", actual, expected)

		}

	}
}
