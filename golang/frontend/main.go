package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
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
	// b3 := b3.New() // nope
	// b3 := b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader | b3.B3SingleHeader)) // nope

	b3 := b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader))
	otel.SetTextMapPropagator(b3)

	tracer = tp.Tracer("greeting-service/year-service")

	mux := http.NewServeMux()
	mux.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		// extract x-request-id from headers and propagate in request
		myHeader := r.Header.Get("x-request-id")

		_, newSpan := tracer.Start(r.Context(), "get ma headers")
		newSpan.SetAttributes(attribute.String("x-request-id", r.Header.Get("x-request-id")))
		newSpan.SetAttributes(attribute.String("x-b3-trace-id", r.Header.Get("x-b3-trace-id")))
		newSpan.SetAttributes(attribute.String("x-b3-spanid", r.Header.Get("x-b3-spanid")))
		newSpan.SetAttributes(attribute.String("x-b3-parentspanid", r.Header.Get("x-b3-parentspanid")))
		newSpan.SetAttributes(attribute.String("x-b3-sampled", r.Header.Get("x-b3-sampled")))
		newSpan.SetAttributes(attribute.String("x-b3-flags", r.Header.Get("x-b3-flags")))
		newSpan.SetAttributes(attribute.String("x-ot-span-context", r.Header.Get("x-ot-span-context")))
		defer newSpan.End()

		name := getName(r.Context(), myHeader)
		message := getMessage(r.Context(), myHeader)

		_, _ = fmt.Fprintf(w, "Hello %s, %s", name, message)
	})

	wrappedHandler := otelhttp.NewHandler(mux, "frontend")

	log.Println("Listening on http://localhost:7000/greeting")
	log.Fatal(http.ListenAndServe(":7000", wrappedHandler))
}

func getName(ctx context.Context, headerId string) string {
	var getNameSpan trace.Span
	ctx, getNameSpan = tracer.Start(ctx, "✨ call /name ✨")
	defer getNameSpan.End()
	return makeRequest(ctx, nameServiceUrl, headerId)
}

func getMessage(ctx context.Context, headerId string) string {
	var getMessageSpan trace.Span
	ctx, getMessageSpan = tracer.Start(ctx, "✨ call /message ✨")
	defer getMessageSpan.End()
	return makeRequest(ctx, messageServiceUrl, headerId)
}

func makeRequest(ctx context.Context, url, headerId string) string {
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	req.Header.Add("x-request-id", headerId)
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
