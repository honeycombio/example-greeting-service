package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	tracer trace.Tracer
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
			semconv.ServiceNameKey.String("message-go"),
		))
	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)
}

func calculateMessage() string {
	messages := []string{
		"how are you?", "how are you doing?", "what's good?", "what's up?", "how do you do?",
		"sup?", "good day to you", "how are things?", "howzit?", "woohoo",
	}
	return messages[rand.Intn(len(messages))]
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	_, newSpan := tracer.Start(r.Context(), "get ma headers")
	newSpan.SetAttributes(attribute.String("x-request-id", r.Header.Get("x-request-id")))
	newSpan.SetAttributes(attribute.String("x-b3-trace-id", r.Header.Get("x-b3-trace-id")))
	newSpan.SetAttributes(attribute.String("x-b3-spanid", r.Header.Get("x-b3-spanid")))
	newSpan.SetAttributes(attribute.String("x-b3-parentspanid", r.Header.Get("x-b3-parentspanid")))
	newSpan.SetAttributes(attribute.String("x-b3-sampled", r.Header.Get("x-b3-sampled")))
	newSpan.SetAttributes(attribute.String("x-b3-flags", r.Header.Get("x-b3-flags")))
	newSpan.SetAttributes(attribute.String("x-ot-span-context", r.Header.Get("x-ot-span-context")))
	defer newSpan.End()

	ctx := r.Context()
	message := func(ctx context.Context) string {
		_, span := tracer.Start(ctx, "look up message")
		defer span.End()
		return calculateMessage()
	}(ctx)

	_, _ = fmt.Fprintf(w, "%v", message)

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
	b3 := b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader))
	otel.SetTextMapPropagator(b3)

	tracer = tp.Tracer("greeting-service/message-service")

	mux := http.NewServeMux()
	mux.HandleFunc("/message", messageHandler)

	wrappedHandler := otelhttp.NewHandler(mux, "message")

	log.Println("Listening on http://localhost:9000/message")
	log.Fatal(http.ListenAndServe(":9000", wrappedHandler))
}
