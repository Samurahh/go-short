package router

import (
	"github.com/Samurahh/go-short/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/shorten-url", controller.ShortenURLHandler).Methods("POST")
	r.HandleFunc("/api/v1/shortened-url/{shortUrl}", controller.ShortenedURLInfoHandler).Methods("GET")
	r.HandleFunc("/api/v1/redirect/{shortURL}", controller.RedirectHandler).Methods("GET")
	return r
}
