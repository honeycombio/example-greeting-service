package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/instrumentation/httptrace"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	years := []int{2015, 2016, 2017, 2018, 2019, 2020}
	exp, err := otlp.NewExporter(
		otlp.WithAddress("localhost:9090"),
		otlp.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}

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
			trace.ContextWithRemoteSpanContext(r.Context(), spanCtx),
			"year-service",
			trace.WithAttributes(attrs...),
		)
		defer span.End()

		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)

		fmt.Fprintf(w, "%d", years[rand.Intn(len(years))])
	})

	log.Fatal(http.ListenAndServe(":6001", mux))
}
