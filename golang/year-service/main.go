package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/contrib/propagators/b3"
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

// func getGrpcEndpoint() string {
// 	apiEndpoint, exists := os.LookupEnv("HONEYCOMB_API_ENDPOINT")
// 	if !exists {
// 		apiEndpoint = "api.honeycomb.io:443"
// 	} else {
// 		u, err := url.Parse(apiEndpoint)
// 		if err != nil {
// 			panic(fmt.Errorf("error %s parsing url: %s", err, apiEndpoint))
// 		}
// 		var host, port string
// 		if u.Port() != "" {
// 			host, port, _ = net.SplitHostPort(u.Host)
// 		} else {
// 			host = u.Host
// 			port = "443"
// 		}
// 		apiEndpoint = fmt.Sprintf("%s:%s", host, port)
// 	}
// 	return apiEndpoint
// }

// func newExporter(ctx context.Context) (*otlptrace.Exporter, error) {
// 	opts := []otlptracegrpc.Option{
// 		otlptracegrpc.WithEndpoint(getGrpcEndpoint()),
// 		otlptracegrpc.WithHeaders(map[string]string{
// 			"x-honeycomb-team":    os.Getenv("HONEYCOMB_API_KEY"),
// 			"x-honeycomb-dataset": os.Getenv("HONEYCOMB_DATASET"),
// 		}),
// 		otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, "")),
// 	}

// 	client := otlptracegrpc.NewClient(opts...)
// 	return otlptrace.New(ctx, client)
// }

// func newTraceProvider(exp *otlptrace.Exporter) *sdktrace.TracerProvider {
// 	r, _ := resource.Merge(
// 		resource.Default(),
// 		resource.NewWithAttributes(
// 			semconv.SchemaURL,
// 			semconv.ServiceNameKey.String("year-go"),
// 		))

// 	return sdktrace.NewTracerProvider(
// 		sdktrace.WithBatcher(exp),
// 		sdktrace.WithResource(r),
// 	)
// }

func newExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	client := otlptracegrpc.NewClient()
	return otlptrace.New(ctx, client)
}

func newTraceProvider(exp *otlptrace.Exporter) *sdktrace.TracerProvider {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("year-go"),
		))
	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)
}

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

	exp, err := newExporter(ctx)
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	tp := newTraceProvider(exp)

	// Handle this error in a sensible manner where possible
	defer func() { _ = tp.Shutdown(ctx) }()

	// Set the Tracer Provider and the W3C Trace Context propagator as globals.
	// Important, otherwise this won't let you see distributed traces be connected!
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
			b3.New()),
	)

	tracer = tp.Tracer("greeting-service/year-service")

	mux := http.NewServeMux()
	mux.HandleFunc("/year", yearHandler)

	wrappedHandler := otelhttp.NewHandler(mux, "year")

	log.Println("Listening on http://localhost:6001/year")
	log.Fatal(http.ListenAndServe(":6001", wrappedHandler))
}
