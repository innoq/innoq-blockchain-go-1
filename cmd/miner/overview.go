package main

import (
	"net/http"

	"encoding/json"

	uuid "github.com/satori/go.uuid"
)

type Overview struct {
	NodeId             string `json:"nodeId"`
	CurrentBlockHeight int64  `json:"currentBlockHeight"`
}

func (o *Overview) serveJson(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(o)
}

func NewOverview() *Overview {
	return &Overview{
		NodeId:             uuid.NewV4().String(),
		CurrentBlockHeight: 1,
	}
}
