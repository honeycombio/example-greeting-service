package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	beeline "github.com/honeycombio/beeline-go"
	"github.com/honeycombio/beeline-go/propagation"
	"github.com/honeycombio/beeline-go/wrappers/config"
	"github.com/honeycombio/beeline-go/wrappers/hnynethttp"
)

func traceParserHook(r *http.Request) *propagation.PropagationContext {
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

func main() {
	beeline.Init(beeline.Config{
		WriteKey:    os.Getenv("HONEYCOMB_API_KEY"),
		Dataset:     os.Getenv("HONEYCOMB_DATASET"),
		ServiceName: "message-go",
	})
	defer beeline.Close()

	messages := []string{
		"how are you?", "how are you doing?", "what's good?", "what's up?", "how do you do?",
		"sup?", "good day to you", "how are things?", "howzit?", "woohoo",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx, span := beeline.StartSpan(ctx, "📖 look up message ✨")
		defer span.Send()
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		fmt.Fprintf(w, messages[rand.Intn(len(messages))])
	})

	log.Fatal(
		http.ListenAndServe(
			":9000",
			hnynethttp.WrapHandlerWithConfig(
				mux,
				config.HTTPIncomingConfig{
					HTTPParserHook: traceParserHook})))
}
