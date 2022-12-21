package token

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateToken(t *testing.T) {
	authId := 1

	token, err := GenerateToken(authId)
	fmt.Println(token)
	require.Nil(t, err)
	require.NotEmpty(t, token)

}

func TestValidateToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7IkF1dGhJZCI6MSwiRXhwaXJlZCI6IjIwMjItMTItMjFUMjM6MjY6MzEuMTE0OTMzKzA3OjAwIn19.toj22OVFDI86apWcXlq0-HquTnnM0f8P_kkIBESN5ok"

	token, err := ValidateToken(tokenString)
	require.Nil(t, err)
	require.NotNil(t, token)
}
