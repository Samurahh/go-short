package generator

import "github.com/golang-jwt/jwt"

var secretKey = []byte("SECRET")

func GenerateToken(url string) (string, error) {
	var claims = jwt.MapClaims{
		"url": url,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := jwtToken.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signed, nil
}
