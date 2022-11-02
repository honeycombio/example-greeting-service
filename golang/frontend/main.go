package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	nameServiceUrl    = os.Getenv("NAME_ENDPOINT") + "/name"
	messageServiceUrl = os.Getenv("MESSAGE_ENDPOINT") + "/message"
	tracer            trace.Tracer
)

func newExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	client := otlptracegrpc.NewClient()
	return otlptrace.New(ctx, client)
}

func newTraceProvider(exp *otlptrace.Exporter) *sdktrace.TracerProvider {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("frontend-go"),
		))
	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)
}

func main() {
	ctx := context.Background()

	exp, err := newExporter(ctx)
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	tp := newTraceProvider(exp)

	// Handle this error in a sensible manner where possible
	defer func() { _ = tp.Shutdown(ctx) }()

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
	)

	tracer = tp.Tracer("greeting-service/year-service")

	mux := http.NewServeMux()
	mux.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {

		name := getName(r.Context())
		message := getMessage(r.Context())

		_, _ = fmt.Fprintf(w, "Hello %s, %s", name, message)
	})

	wrappedHandler := otelhttp.NewHandler(mux, "frontend")

	log.Println("Listening on http://localhost:7000/greeting")
	log.Fatal(http.ListenAndServe(":7000", wrappedHandler))
}

func getName(ctx context.Context) string {
	var getNameSpan trace.Span
	ctx, getNameSpan = tracer.Start(ctx, "✨ call /name ✨")
	defer getNameSpan.End()
	return makeRequest(ctx, nameServiceUrl)
}

func getMessage(ctx context.Context) string {
	var getMessageSpan trace.Span
	ctx, getMessageSpan = tracer.Start(ctx, "✨ call /message ✨")
	defer getMessageSpan.End()
	return makeRequest(ctx, messageServiceUrl)
}

func makeRequest(ctx context.Context, url string) string {
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(body)
}
