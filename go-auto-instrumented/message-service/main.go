package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	messages := []string{
		"how are you?", "how are you doing?", "what's good?", "what's up?", "how do you do?",
		"sup?", "good day to you", "how are things?", "howzit?", "woohoo",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		_, _ = fmt.Fprintf(w, messages[rand.Intn(len(messages))])
	})
   	handler := http.Handler(mux)

	log.Println("Listening on http://localhost:9000/message")
	log.Fatal(http.ListenAndServe(":9000",handler))
}
