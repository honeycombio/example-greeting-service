package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/honeycombio/honeycomb-opentelemetry-go"

	"github.com/honeycombio/otel-config-go/otelconfig"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
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
	otelShutdown, err := otelconfig.ConfigureOpenTelemetry()
	if err != nil {
		log.Fatalf("error setting up OTel SDK - %e", err)
	}
	defer otelShutdown()

	tracer = otel.Tracer("greeting-service/year-service")

	mux := http.NewServeMux()
	mux.HandleFunc("/year", yearHandler)

	wrappedHandler := otelhttp.NewHandler(mux, "year")

	log.Println("Listening on http://localhost:6001/year")
	log.Fatal(http.ListenAndServe(":6001", wrappedHandler))
}
