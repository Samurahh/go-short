package controller

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/shorten-url", ShortenURLHandler).Methods("POST")
	r.HandleFunc("/api/v1/shortened-url/{shortUrl}", ShortenedURLInfoHandler).Methods("GET")
	r.HandleFunc("/api/v1/redirect/{shortURL}", RedirectHandler).Methods("GET")
	return r
}
