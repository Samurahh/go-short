package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/Samurahh/go-short/generator"
)

var httpServer http.Server = http.Server{
	Addr: ":8080", // default port
}

func main() {
	http.HandleFunc("/gen", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, html.EscapeString(generator.ShortUrl()))
		w.Header().Add("Content-Type", "text/plain")
	})
	log.Fatal(httpServer.ListenAndServe())
}
