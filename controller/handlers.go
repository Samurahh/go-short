package controller

import (
	"log"
	"net/http"
)

// @Tag User
// @Title Create Short URL - API
// @Description Creates an expirable short URL for the given URL
// @Param request body model.CreateShortURLRequest
// @Success 200 (object) model.ShortenedURLResponse Shortened URL response
// @Router /api/v1/shorten-url [post]
func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/api/v1/shorten-url")
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
