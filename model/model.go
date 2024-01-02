package model

import "time"

type CreateShortURLRequest struct {
	Url            string     `json:"url"`
	ExpirationTime *time.Time `json:"expirationTime"`
	Id             string     `json:"id"`
}

type ShortenedURL struct {
	OriginalUrl    string     `json:"originalUrl"`
	ShortUrl       string     `json:"shortUrl"`
	ExpirationTime *time.Time `json:"expirationTime"`
	AddedTime      time.Time  `json:"addedTime"`
	OwnershipToken string     `json:"ownershipToken"`
}

type ShortenedURLResponse struct {
	ShortenedURL
	Timestamp time.Time `json:"timestamp"`
}
