package generator

import "github.com/golang-jwt/jwt"

func GenerateToken(url string) (string, error) {
	var claims = jwt.MapClaims{
		"url": url,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signed, err := jwtToken.SignedString("SECRET")
	if err != nil {
		return "", err
	}

	return signed, nil
}
