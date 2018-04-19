package main

type Block struct {
	Index             string   `json:"index"`
	Timestamp         string   `json:"timestamp"`
	Proof             uint64   `json:"proof"`
	Transaction       []string `json:"transaction"`
	PreviousBlockHash string   `json:"previousBlockHash"`
}
