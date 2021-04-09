package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/correlation"
	"go.opentelemetry.io/otel/instrumentation/httptrace"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/honeycombio/opentelemetry-exporter-go/honeycomb"
)

const (
	nameServiceUrl = "http://localhost:8000/name"
	messageServiceUrl = "http://localhost:9000/message"
)

func main() {
	exp, err := honeycomb.NewExporter(
		honeycomb.Config{
			APIKey: os.Getenv("HONEYCOMB_WRITE_KEY"),
		},
		honeycomb.TargetingDataset(os.Getenv("HONEYCOMB_DATASET")),
		honeycomb.WithServiceName("frontend"),
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

	tracer := global.Tracer("greeting-service/frontend")

	mux := http.NewServeMux()
	mux.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		attrs, entries, spanCtx := httptrace.Extract(r.Context(), r)

		r = r.WithContext(correlation.ContextWithMap(r.Context(), correlation.NewMap(correlation.MapUpdate{MultiKV: entries})))

		ctx, span := tracer.Start(
			trace.ContextWithRemoteSpanContext(r.Context(), spanCtx),
			"greeting",
			trace.WithAttributes(attrs...),
		)
		defer span.End()

		name := getName(ctx)
		message := getMessage(ctx)

		fmt.Fprintf(w, "Hello %s, %s", name, message)
	})

	log.Fatal(http.ListenAndServe(":7000", mux))
}

func getName(ctx context.Context) string {
	tracer := global.Tracer("greeting-service/frontend")
	var getNameSpan trace.Span
  ctx, getNameSpan = tracer.Start(ctx, "name retrieval")
  defer getNameSpan.End()
	return makeRequest(ctx, nameServiceUrl)
}

func getMessage(ctx context.Context) string {
	tracer := global.Tracer("greeting-service/frontend")
	var getMessageSpan trace.Span
  ctx, getMessageSpan = tracer.Start(ctx, "message retrieval")
  defer getMessageSpan.End()
	return makeRequest(ctx, messageServiceUrl)
}

func makeRequest(ctx context.Context, url string) string {
	req, _ := http.NewRequest("GET", url, nil)
	ctx, req = httptrace.W3C(ctx, req)
	httptrace.Inject(ctx, req)
	client := http.DefaultClient
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

