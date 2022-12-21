package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	TOKEN_Key    = "FSUnIQPYBvUeUbo6loiJp4HVk515eA7i"
	TOKEN_Expiry = 10 * 60 * time.Second
)

func GenerateToken(authId int) (string, error) {
	payload := Token{
		AuthId:  authId,
		Expired: time.Now().Add(TOKEN_Expiry),
	}
	claims := jwt.MapClaims{
		"payload": payload,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(TOKEN_Key))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateToken(tokenString string) (*Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(TOKEN_Key), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		payloadInterface := claims["payload"]

		payload := Token{}

		payloadByte, err := json.Marshal(payloadInterface)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(payloadByte, &payload)
		if err != nil {
			return nil, err
		}

		now := time.Now()
		if now.After(payload.Expired) {
			return nil, errors.New("token expired")
		}
		return &payload, nil
	} else {
		return nil, errors.New("token invalid")
	}
}
