package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "supersecret"

func GenerateToken(email string, id int, role string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"role":  role,
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

type MyCustomClaims struct {
	Email string          `json:"email"`
	ID    int             `json:"id"`
	NBF   jwt.NumericDate `json:"nbf"`
}

// map[email:hongnhu@gmail.com id:5 nbf:1.7287344e+09]

func TokenCheck(jwtToken string) (jwt.MapClaims, error) {

	token, err := parseToken(jwtToken)

	if err != nil {
		return nil, err
	}

	data, Ok := token.Claims.(jwt.MapClaims)

	if !Ok {
		return nil, errors.New("unable to map claims")
	}

	return data, nil
}
