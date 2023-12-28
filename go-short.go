package main

import (
	"log"
	"net/http"

	"github.com/Samurahh/go-short/controller"
)

var httpServer *http.Server = &http.Server{
	Addr:    ":8080", // default port
	Handler: controller.Router(),
}

func main() {
	log.Fatal(httpServer.ListenAndServe())
}
