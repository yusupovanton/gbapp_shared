package lib_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yusupovanton/gbapp_shared"
)

func TestDbActions(t *testing.T) {

	t.Parallel()
	
	db, err := lib.ConnectToDB()

	require.NotNil(t, db)
	require.NoError(t, err)
}
