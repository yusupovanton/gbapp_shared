package db_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	
	. "shared/db"
)

func TestDbActions(t *testing.T) {

	t.Parallel()
	
	db, err := ConnectToDB()

	require.NotNil(t, db)
	require.NoError(t, err)
}
