// +build integration

package sqlstore

import (
	"testing"

	"github.com/grafarg/grafarg/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestGetDBHealthQuery(t *testing.T) {
	InitTestDB(t)

	query := models.GetDBHealthQuery{}
	err := GetDBHealthQuery(&query)
	require.NoError(t, err)
}
