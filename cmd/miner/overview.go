package main

import (
	"net/http"

	"encoding/json"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	uuid "github.com/satori/go.uuid"
)

type Overview struct {
	NodeId             string `json:"nodeId"`
	CurrentBlockHeight uint64 `json:"currentBlockHeight"`
	chain              *Chain
}

func (o *Overview) serveJson(w http.ResponseWriter, r *http.Request) {

	wireContext, _ := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))

	span := opentracing.StartSpan("Overview:serveJson", ext.RPCServerOption(wireContext))
	defer span.Finish()

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
