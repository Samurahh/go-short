package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Samurahh/go-short/controller"
)

// By default it listens to Addr: ":8080" as configured in main()
var httpServer *http.Server = &http.Server{
	Handler: controller.Router(),
}

func main() {
	var port int
	var address string
	flag.IntVar(&port, "p", 8080, "port to run the http server on | default 8080")
	flag.IntVar(&port, "port", 8080, "port to run the http server on | default 8080")
	flag.StringVar(&address, "address", "", "address the http server will run on | default 127.0.0.1 / localhost")
	flag.Parse()

	if port < 0 || port > 65535 {
		log.Panicf("Port %v out of range [0, 65535]", port)
	}
	httpServer.Addr = fmt.Sprintf("%s:%v", address, port)

	log.Fatal(httpServer.ListenAndServe())
}
