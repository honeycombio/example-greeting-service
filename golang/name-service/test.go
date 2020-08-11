package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"crypto/tls"

	beeline "github.com/honeycombio/beeline-go"
	"github.com/honeycombio/beeline-go/propagation"
	"github.com/honeycombio/beeline-go/wrappers/config"
	"github.com/honeycombio/beeline-go/wrappers/hnynethttp"
)

func main() {
	beeline.Init(beeline.Config{
		WriteKey: os.Getenv("HONEYCOMB_WRITE_KEY"),
		Dataset: os.Getenv("HONEYCOMB_DATASET"),
		ServiceName: "name-service-golang",
    })
    defer beeline.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		req, _ := http.NewRequest("GET", "https://service-alb-tf-1313179747.us-east-1.elb.amazonaws.com/foobar", nil)
		req = req.WithContext(r.Context())
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{
			Transport: hnynethttp.WrapRoundTripperWithConfig(transport, config.HTTPOutgoingConfig{
				HTTPPropagationHook: propagateTraceHook,
			}),
			Timeout: time.Second * 5,
		}
		fmt.Println("Calling service...")
		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(body))
	})

	log.Println("Listening on ", ":8000")
	log.Fatal(http.ListenAndServe(":8000", hnynethttp.WrapHandler(mux)))
}

func propagateTraceHook(r *http.Request, prop *propagation.PropagationContext) map[string]string {
	fmt.Println("Hi, I'm in a propagate hook")
	fmt.Printf("%+v\n", prop)
	//ctx := r.Context()
	//ctx, headers := propagation.MarshalW3CTraceContext(ctx, prop)
	headers := propagation.MarshalAmazonTraceContext(prop)
	fmt.Printf("%+v\n", headers)
	return map[string]string{
		"X-Amzn-Trace-Id": headers,
	}
}
