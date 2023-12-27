package main

import (
	"log"
	"net/http"

	"github.com/Samurahh/go-short/router"
)

var httpServer *http.Server = &http.Server{
	Addr: ":8080", // default port
	Handler: router.Router(),
}

func main() {
	log.Fatal(httpServer.ListenAndServe())
}
