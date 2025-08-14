package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var (
	tracer trace.Tracer
)

func calculateYear() int {
	years := []int{2016, 2017, 2018, 2019, 2020}
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
	return years[rand.Intn(len(years))]
}

func yearHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	year := func(ctx context.Context) int {
		_, span := tracer.Start(ctx, "calculate-year")
		defer span.End()

		return calculateYear()
	}(ctx)

	_, _ = fmt.Fprintf(w, "%d", year)
}

func main() {
	ctx := context.Background()
	exporter, _ := otlptrace.New(ctx, otlptracegrpc.NewClient(
		otlptracegrpc.WithEndpoint(os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")),
		otlptracegrpc.WithHeaders(map[string]string{
			"x-honeycomb-team": os.Getenv("HONEYCOMB_API_KEY"),
		}),
	))
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
	)
	defer tracerProvider.Shutdown(ctx)

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
	)

	tracer = otel.Tracer("greeting-service/year-service")

	mux := http.NewServeMux()
	mux.HandleFunc("/year", yearHandler)

	wrappedHandler := otelhttp.NewHandler(mux, "year")

	log.Println("Listening on http://localhost:6001/year")
	log.Fatal(http.ListenAndServe(":6001", wrappedHandler))
}
