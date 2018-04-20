package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	thislog "github.com/innoq-blockchain-go-1/pkg/log"
	"github.com/innoq-blockchain-go-1/pkg/tracing"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-lib/metrics/go-kit"
	"github.com/uber/jaeger-lib/metrics/go-kit/expvar"
	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewDevelopment()
	logfac := thislog.NewFactory(logger.With(zap.String("service", "miner")))
	metricsFactory := xkit.Wrap("", expvar.NewFactory(10))
	opentracing.InitGlobalTracer(tracing.Init("Miner", metricsFactory.Namespace("miner", nil), logfac))

	r := mux.NewRouter()

	events := NewEvents()
	events.Start()
	defer events.Stop()

	chain := NewChain()
	miner := NewMiner(chain, events, "00000")
	overview := NewOverview(chain)
	transactions := NewTransactions(*events)

	miner.Start()
	defer miner.Stop()

	r.HandleFunc("/", overview.serveJson)

	r.HandleFunc("/mine", miner.mine)

	r.HandleFunc("/blocks", chain.serveJson)

	r.Handle("/events", events)

	r.HandleFunc("/transactions", transactions.Post)

	r.HandleFunc("/transactions/{id}", transactions.serveJson)

	r.HandleFunc("/ui", GetIndex)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
