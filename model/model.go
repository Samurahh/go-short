package model

import "time"

type CreateShortURLRequest struct {
	Url            string    `json:"url"`
	ExpirationTime *time.Time `json:"expirationTime"`
}

type ShortenedURL struct {
	OriginalUrl    string    `json:"originalUrl"`
	ShortUrl       string    `json:"shortUrl"`
	ExpirationTime *time.Time `json:"expirationTime"`
	AddedTime      time.Time `json:"addedTime"`
}

type ShortenedURLResponse struct {
	ShortenedURL
	Timestamp time.Time `json:"timestamp"`
}
