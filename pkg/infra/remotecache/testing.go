package remotecache

import (
	"testing"

	"github.com/famarker/grafarg/pkg/services/sqlstore"
	"github.com/famarker/grafarg/pkg/setting"
	"github.com/stretchr/testify/require"
)

// NewFakeStore creates store for testing
func NewFakeStore(t *testing.T) *RemoteCache {
	t.Helper()

	opts := &setting.RemoteCacheOptions{
		Name:    "database",
		ConnStr: "",
	}

	SQLStore := sqlstore.InitTestDB(t)

	dc := &RemoteCache{
		SQLStore: SQLStore,
		Cfg: &setting.Cfg{
			RemoteCacheOptions: opts,
		},
	}

	err := dc.Init()
	require.NoError(t, err, "Failed to init remote cache for test")

	return dc
}
