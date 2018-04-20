package main

import "testing"

func TestNewChain(t *testing.T) {
	chain := NewChain()
	height := chain.Height()
	if height != 1 {
		t.Errorf("NewChain(): Unexpected block height, got %d, want %d.", height, 1)
	}
}

func TestLastBlock(t *testing.T) {
	chain := NewChain()
	lastBlock := chain.LastBlock()
	if lastBlock.Index != 1 {
		t.Errorf("Chain#LastBlock(): First block should have index 1, found %d.", lastBlock.Index)
	}
}

func TestBlocks(t *testing.T) {
	chain := NewChain()
	blocks := chain.Blocks()
	if len(blocks) != 1 {
		t.Errorf("Chain#Blocks(): Didn't return blocks.")
	}
}

func TestAddBlock(t *testing.T) {
	chain := NewChain()
	chain.addBlock(Block{})
	height := chain.Height()
	if height != 2 {
		t.Errorf("Chain#addBlock(): Slice didn't grow.")
	}
}

// testing findTransactionById (our blockchain query function)

func TestFindInEmptyChain(t *testing.T) {
	chain := NewChain()

	transaction := findTransactionById(chain, "42")

	if transaction != nil {
		t.Errorf("Block#findTransactionById(): found transaction in empty chain")
	}

	// try to find magicGenesisTransactionID
	transaction = findTransactionById(chain, magicGenesisTransactionID)

	if transaction == nil {
		t.Errorf("%s not found in GenesisBlock", magicGenesisTransactionID)
	}

}
