package encryption

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGeneratePassword(t *testing.T) {
	passwordString := "IniPassword"
	pass, err := HashPassword(passwordString)
	fmt.Println(pass)
	require.Nil(t, err)
	require.NotEmpty(t, pass)
}

func TestValidatePassword(t *testing.T) {
	passwordString := "IniPassword"
	passwordHash := "$2a$10$lfT876pBo9IndYvuK3C/w.IKgLx3/k8pYam7lMZOyV2sjyu2luaxi"

	err := ValidatePassword(passwordHash, passwordString)
	require.Nil(t, err)
}
