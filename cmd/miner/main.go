package main

import (
	"log"
	"net/http"

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

	events := NewEvents()
	events.Start()
	defer events.Stop()

	chain := NewChain()
	miner := NewMiner(chain, events, "00000")
	overview := NewOverview(chain)

	miner.Start()
	defer miner.Stop()

	http.HandleFunc("/", overview.serveJson)

	http.HandleFunc("/mine", miner.mine)

	http.HandleFunc("/blocks", chain.serveJson)

	http.Handle("/events", events)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
