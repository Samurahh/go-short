package generator

import "github.com/golang-jwt/jwt"

func GenerateToken(url string) *jwt.Token {
	var claims = jwt.MapClaims{
		"url":    url,
		"secret": "SECRET",
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return jwtToken
}
