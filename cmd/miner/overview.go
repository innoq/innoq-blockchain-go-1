package main

import (
	"net/http"

	"encoding/json"

	uuid "github.com/satori/go.uuid"
)

type Overview struct {
	NodeId             string `json:"nodeId"`
	CurrentBlockHeight uint64 `json:"currentBlockHeight"`
	chain              *Chain
}

func (o *Overview) serveJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	o.CurrentBlockHeight = o.chain.Height()
	json.NewEncoder(w).Encode(o)
}

func NewOverview(chain *Chain) *Overview {
	return &Overview{
		NodeId:             uuid.NewV4().String(),
		CurrentBlockHeight: 1,
		chain:              chain,
	}
}
