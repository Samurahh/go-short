package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Samurahh/go-short/generator"
	"github.com/Samurahh/go-short/model"
	"github.com/google/uuid"
)

// @Tag User
// @Title Create Short URL - API
// @Description Creates an expirable short URL for the given URL
// @Param request body model.CreateShortURLRequest
// @Success 200 (object) model.ShortenedURLResponse Shortened URL response
// @Router /api/v1/shorten-url [post]
func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/api/v1/shorten-url")
	var CreateShortURLRequest []model.CreateShortURLRequest

	err := json.NewDecoder(r.Body).Decode(&CreateShortURLRequest)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var ReturnShortURLs []model.ShortenedURLResponse

	for _, v := range CreateShortURLRequest {
		if v.Url == "" {
			http.Error(w, "Url cannot be empty", http.StatusBadRequest)
			return
		}
		uuid, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		v.Id = uuid.String()

		token, err := generator.GenerateToken(v.Url)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var shortenedURL = model.ShortenedURL{
			OriginalUrl:    v.Url,
			ShortUrl:       generator.ShortUrl(),
			ExpirationTime: v.ExpirationTime,
			AddedTime:      time.Now(),
			OwnershipToken: token,
		}

		ReturnShortURLs = append(ReturnShortURLs, model.ShortenedURLResponse{
			ShortenedURL: shortenedURL,
			Timestamp:    time.Now(),
		})
	}

	err = json.NewEncoder(w).Encode(ReturnShortURLs)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Tag User
// @Title Retrieve Info on Short URL - API
// @Description Retrieves information about the short URL, when it was created, what it points to and when it will expire.
// @Param shortUrl path string
// @Success 200 (object) model.ShortenedURL Info about the short URL
// @Router /api/v1/shortened-url/{shortUrl} [get]
func ShortenedURLInfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/api/v1/shortened-url/{shortUrl}")
}

// @Tag User
// @Title Redirect to Short URL destination - API
// @Description Responds with the destination URL for the Short URL specified.
// @Param shortURL path string
// @Success 302 (object) string Destination of Short URL
// @Router /api/v1/redirect/{shortURL}
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/api/v1/redirect/{shortURL}")
}
