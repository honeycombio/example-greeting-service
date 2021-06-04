package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/instrumentation/httptrace"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/honeycombio/opentelemetry-exporter-go/honeycomb"
)

func main() {
	years := []int{2015, 2016, 2017, 2018, 2019, 2020}
	exp, err := honeycomb.NewExporter(
		honeycomb.Config{
			APIKey: os.Getenv("HONEYCOMB_API_KEY"),
		},
		honeycomb.TargetingDataset(os.Getenv("HONEYCOMB_DATASET")),
		honeycomb.WithServiceName("year-go"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer exp.Close()

	config := sdktrace.Config{
		DefaultSampler: sdktrace.AlwaysSample(),
	}
	tp, err := sdktrace.NewProvider(sdktrace.WithConfig(config), sdktrace.WithSyncer(exp))
	if err != nil {
		log.Fatal(err)
	}
	global.SetTraceProvider(tp)

	tracer := global.Tracer("greeting-service/year-service")

	mux := http.NewServeMux()

	mux.HandleFunc("/year", func(w http.ResponseWriter, r *http.Request) {
		attrs, _, spanCtx := httptrace.Extract(r.Context(), r)
		_, span := tracer.Start(
			trace.ContextWithRemoteSpanContext(r.Context(), spanCtx), "/year",
			trace.WithAttributes(attrs...),
		)
		defer span.End()

		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)

		fmt.Fprintf(w, "%d", years[rand.Intn(len(years))])
	})

	log.Fatal(http.ListenAndServe(":6001", mux))
}
