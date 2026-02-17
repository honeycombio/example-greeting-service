package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/contrib/otelconf"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var (
	yearServiceUrl = os.Getenv("YEAR_ENDPOINT") + "/year"
	tracer         trace.Tracer
)

func main() {
	// initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	sdk, err := otelconf.NewSDK()
	if err != nil {
		log.Fatal(err)
	}
	defer sdk.Shutdown(context.Background())

	otel.SetTracerProvider(sdk.TracerProvider())
	otel.SetTextMapPropagator(sdk.Propagator())

	tracer = otel.Tracer("greeting-service/year-service")

	namesByYear := map[int][]string{
		2016: {"sophia", "jackson", "emma", "aiden", "olivia", "lucas", "ava", "liam", "mia", "noah"},
		2017: {"sophia", "jackson", "olivia", "liam", "emma", "noah", "ava", "aiden", "isabella", "lucas"},
		2018: {"sophia", "jackson", "olivia", "liam", "emma", "noah", "ava", "aiden", "isabella", "caden"},
		2019: {"sophia", "liam", "olivia", "jackson", "emma", "noah", "ava", "aiden", "aria", "grayson"},
		2020: {"olivia", "noah", "emma", "liam", "ava", "elijah", "isabella", "oliver", "sophia", "lucas"},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		year, _ := getYear(r.Context())
		names := namesByYear[year]
		_, _ = fmt.Fprintf(w, names[rand.Intn(len(names))])
	})

	wrappedHandler := otelhttp.NewHandler(mux, "name")

	log.Println("Listening on http://localhost:8000/name")
	log.Fatal(http.ListenAndServe(":8000", wrappedHandler))
}

func getYear(ctx context.Context) (int, context.Context) {
	ctx, span := tracer.Start(ctx, "✨ call /year ✨")
	defer span.End()
	req, err := http.NewRequestWithContext(ctx, "GET", yearServiceUrl, nil)
	if err != nil {
		fmt.Printf("error creating request: %s", err)
	}
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
	year, err := strconv.Atoi(string(body))
	if err != nil {
		panic(err)
	}
	return year, ctx
}
