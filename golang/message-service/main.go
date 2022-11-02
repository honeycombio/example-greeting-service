package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
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

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
	)

	tracer = tp.Tracer("greeting-service/message-service")

	mux := http.NewServeMux()
	mux.HandleFunc("/message", messageHandler)

	wrappedHandler := otelhttp.NewHandler(mux, "message")

	log.Println("Listening on http://localhost:9000/message")
	log.Fatal(http.ListenAndServe(":9000", wrappedHandler))
}
