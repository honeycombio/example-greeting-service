package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func calculateYear() int {
	years := []int{2016, 2017, 2018, 2019, 2020}
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
	return years[rand.Intn(len(years))]
}

func yearHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	year := func(ctx context.Context) int {
		return calculateYear()
	}(ctx)

	_, _ = fmt.Fprintf(w, "%d", year)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/year", yearHandler)

	handler := http.Handler(mux)

	log.Println("Listening on http://localhost:6001/year")
	log.Fatal(http.ListenAndServe(":6001", handler))
}
