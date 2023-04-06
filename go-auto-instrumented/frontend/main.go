package main

import (
	"context"
	"fmt"
    "net/http"
	"io/ioutil"
	"log"
	"os"
)

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

var (
	nameServiceUrl    = getEnv("NAME_ENDPOINT", "http://localhost:8000") + "/name"
	messageServiceUrl = getEnv("MESSAGE_ENDPOINT", "http://localhost:9000") + "/message"
)

func main() {

    mux := http.NewServeMux()
	mux.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		name := getName(r.Context())
		message := getMessage(r.Context())

		_, _ = fmt.Fprintf(w, "Hello %s, %s", name, message)
	})

	handler := http.Handler(mux)

	log.Println("Listening on http://localhost:7007/greeting")
	log.Fatal(http.ListenAndServe(":7007", handler))
}

func getName(ctx context.Context) string {
	return makeRequest(ctx, nameServiceUrl)
}

func getMessage(ctx context.Context) string {
	return makeRequest(ctx, messageServiceUrl)
}

func makeRequest(ctx context.Context, url string) string {
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	client := http.Client{Transport: http.DefaultTransport}
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
