package main

type Transactions struct {
	Id        string `json:"id"`
	Payload   string `json:"payload"`
	Timestamp uint64 `json:"timestamp"`
}
type Block struct {
	Index             uint64         `json:"index"`
	PreviousBlockHash string         `json:"previousBlockHash"`
	Proof             uint64         `json:"proof"`
	Timestamp         uint64         `json:"timestamp"`
	Transactions      []Transactions `json:"transactions"`
}
