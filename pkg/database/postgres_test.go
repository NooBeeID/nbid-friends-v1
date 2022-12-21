package database

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnectDB_Success(t *testing.T) {
	db, err := ConnectDB()

	require.Nil(t, err)
	require.NotNil(t, db)
}
