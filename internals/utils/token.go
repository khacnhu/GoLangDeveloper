package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "supersecret"

func GenerateToken(email string, id int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"nbf":   time.Date(2024, 10, 12, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(secret))

	return tokenStr, err

}

func parseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(tokenStr *jwt.Token) (interface{}, error) {
		if _, ok := tokenStr.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil

}

func TokenCheck(jwtToken string) (interface{}, error) {

	token, err := parseToken(jwtToken)

	if err != nil {
		return nil, err
	}

	data, Ok := token.Claims.(jwt.MapClaims)
	// fmt.Println("email ", data.email)

	if !Ok {
		return nil, errors.New("unable to map claims")
	}

	return data, nil
}
