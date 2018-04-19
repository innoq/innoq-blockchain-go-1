package main

type Transactions struct {
	Id        string `json:"id"`
	Timestamp uint64 `json:"timestamp"`
	Payload   string `json:"payload"`
}
type Block struct {
	Index             uint64         `json:"index"`
	Timestamp         uint64         `json:"timestamp"`
	Proof             uint64         `json:"proof"`
	Transactions      []Transactions `json:"transactions"`
	PreviousBlockHash string         `json:"previousBlockHash"`
}
