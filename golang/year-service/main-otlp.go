package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	ctx := context.Background()
	years := []int{2015, 2016, 2017, 2018, 2019, 2020}
	exp, err := otlp.NewExporter(
		ctx,
		otlp.WithInsecure(),  // comment this out when sending to a TLS endpoint
		otlp.WithAddress(os.Getenv("HONEYCOMB_OTLP_ADDRESS")),
		otlp.WithHeaders(map[string]string{
			"x-honeycomb-team": os.Getenv("HONEYCOMB_WRITE_KEY"),
			"x-honeycomb-dataset": "test-otlp",
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	config := sdktrace.Config{
		DefaultSampler: sdktrace.AlwaysSample(),
	}
	tp := sdktrace.NewTracerProvider(sdktrace.WithConfig(config), sdktrace.WithSyncer(exp))
	otel.SetTracerProvider(tp)

	tracer := otel.Tracer("greeting-service/year-service")


	mux := http.NewServeMux()
	mux.HandleFunc("/year", func(w http.ResponseWriter, r *http.Request) {
		attrs, _, spanCtx := otelhttptrace.Extract(r.Context(), r)
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
