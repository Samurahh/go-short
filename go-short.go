package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/Samurahh/go-short/router"
)

var httpServer *http.Server = &http.Server{
	Addr:    ":8080", // default port
	Handler: router.Router(),
}

func main() {
	var port int
	flag.IntVar(&port, "p", 8080, "port to run the server on")
	flag.Parse()

	if port != 0 {
		portStr := strconv.Itoa(port)
		httpServer.Addr = ":" + portStr
	}

	log.Fatal(httpServer.ListenAndServe())
}
