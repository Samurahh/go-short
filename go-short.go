package main

import (
	"flag"
	"fmt"
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
	var host string
	flag.IntVar(&port, "p", 8080, "port to run the http server on")
	flag.StringVar(&host, "h", "", "hostname the http server will run on")
	flag.Parse()

	if port != 0 {
		portStr := strconv.Itoa(port)
		httpServer.Addr = fmt.Sprintf("%s:%s", host, portStr)
	}

	log.Fatal(httpServer.ListenAndServe())
}
