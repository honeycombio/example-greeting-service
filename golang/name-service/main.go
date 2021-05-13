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

	beeline "github.com/honeycombio/beeline-go"
	"github.com/honeycombio/beeline-go/propagation"
	"github.com/honeycombio/beeline-go/wrappers/config"
	"github.com/honeycombio/beeline-go/wrappers/hnynethttp"

	"go.opentelemetry.io/otel/instrumentation/httptrace"
)

func main() {
	beeline.Init(beeline.Config{
		WriteKey:    os.Getenv("HONEYCOMB_API_KEY"),
		Dataset:     os.Getenv("HONEYCOMB_DATASET"),
		ServiceName: "name-go",
	})
	defer beeline.Close()

	namesByYear := map[int][]string{
		2015: []string{"sophia", "jackson", "emma", "aiden", "olivia", "liam", "ava", "lucas", "mia", "noah"},
		2016: []string{"sophia", "jackson", "emma", "aiden", "olivia", "lucas", "ava", "liam", "mia", "noah"},
		2017: []string{"sophia", "jackson", "olivia", "liam", "emma", "noah", "ava", "aiden", "isabella", "lucas"},
		2018: []string{"sophia", "jackson", "olivia", "liam", "emma", "noah", "ava", "aiden", "isabella", "caden"},
		2019: []string{"sophia", "liam", "olivia", "jackson", "emma", "noah", "ava", "aiden", "aria", "grayson"},
		2020: []string{"olivia", "noah", "emma", "liam", "ava", "elijah", "isabella", "oliver", "sophia", "lucas"},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%+v\n", r.Header)
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		year, _ := getYear(r.Context())
		names := namesByYear[year]
		fmt.Fprintf(w, names[rand.Intn(len(names))])
	})

	traceHeaderParserHook := func(r *http.Request) *propagation.PropagationContext {
		headers := map[string]string{
			"traceparent": r.Header.Get("traceparent"),
		}
		ctx := r.Context()
		ctx, prop, err := propagation.UnmarshalW3CTraceContext(ctx, headers)
		if err != nil {
			fmt.Println("Error unmarshaling header")
			fmt.Println(err)
		}
		return prop
	}

	log.Println("Listening on ", ":8000")
	log.Fatal(http.ListenAndServe(":8000", hnynethttp.WrapHandlerWithConfig(mux, config.HTTPIncomingConfig{HTTPParserHook: traceHeaderParserHook})))
}

func propagateTraceHook(r *http.Request, prop *propagation.PropagationContext) map[string]string {
	ctx := r.Context()
	ctx, headers := propagation.MarshalW3CTraceContext(ctx, prop)
	return headers
}

func getYear(ctx context.Context) (int, context.Context) {
	ctx, span := beeline.StartSpan(ctx, "✨ call /year ✨")
	defer span.Send()
	req, _ := http.NewRequest("GET", "http://localhost:6001/year", nil)
	ctx, req = httptrace.W3C(ctx, req)
	httptrace.Inject(ctx, req)
	client := &http.Client{
		Transport: hnynethttp.WrapRoundTripperWithConfig(http.DefaultTransport, config.HTTPOutgoingConfig{HTTPPropagationHook: propagateTraceHook}),
		Timeout:   time.Second * 5,
	}
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
