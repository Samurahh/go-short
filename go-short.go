package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		timeBytes, _ := time.Now().MarshalText()
		fmt.Fprintf(w, "Hello, world! Time is %v\n", html.EscapeString(string(timeBytes)))
	})

	s := &http.Server{
		Addr: ":6090",
	}

	log.Fatal(s.ListenAndServe())
}
