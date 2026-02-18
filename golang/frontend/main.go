package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/contrib/otelconf"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var (
	nameServiceUrl    = os.Getenv("NAME_ENDPOINT") + "/name"
	messageServiceUrl = os.Getenv("MESSAGE_ENDPOINT") + "/message"
	tracer            trace.Tracer
)

func main() {
	sdk, err := otelconf.NewSDK()
	if err != nil {
		log.Fatal(err)
	}
	defer sdk.Shutdown(context.Background())

	// Set the Tracer Provider and the W3C Trace Context propagator as globals.
	// Important, otherwise this won't let you see distributed traces be connected!
	otel.SetTracerProvider(sdk.TracerProvider())
	otel.SetTextMapPropagator(sdk.Propagator())

	tracer = otel.Tracer("greeting-service/year-service")

	mux := http.NewServeMux()
	mux.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		name := getName(r.Context())
		message := getMessage(r.Context())

		_, _ = fmt.Fprintf(w, "Hello %s, %s", name, message)
	})

	wrappedHandler := otelhttp.NewHandler(mux, "frontend")

	log.Println("Listening on http://localhost:7777/greeting")
	log.Fatal(http.ListenAndServe(":7777", wrappedHandler))
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
