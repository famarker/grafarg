package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/famarker/grafarg/pkg/plugins"
	"github.com/famarker/grafarg/pkg/services/rendering"

	"github.com/famarker/grafarg/pkg/services/licensing"

	"github.com/famarker/grafarg/pkg/bus"
	"github.com/famarker/grafarg/pkg/services/sqlstore"

	"gopkg.in/macaron.v1"

	"github.com/famarker/grafarg/pkg/setting"
)

func setupTestEnvironment(t *testing.T, cfg *setting.Cfg) (*macaron.Macaron, *HTTPServer) {
	t.Helper()
	sqlstore.InitTestDB(t)

	{
		oldVersion := setting.BuildVersion
		oldCommit := setting.BuildCommit
		oldEnv := setting.Env
		t.Cleanup(func() {
			setting.BuildVersion = oldVersion
			setting.BuildCommit = oldCommit
			setting.Env = oldEnv
		})
	}

	bus.ClearBusHandlers()
	bus.AddHandler("sql", sqlstore.GetPluginSettings)
	t.Cleanup(bus.ClearBusHandlers)

	r := &rendering.RenderingService{Cfg: cfg}

	hs := &HTTPServer{
		Cfg:           cfg,
		Bus:           bus.GetBus(),
		License:       &licensing.OSSLicensingService{Cfg: cfg},
		RenderService: r,
		PluginManager: &plugins.PluginManager{Cfg: cfg},
	}

	m := macaron.New()
	m.Use(getContextHandler(t, cfg).Middleware)
	m.Use(macaron.Renderer(macaron.RenderOptions{
		Directory:  filepath.Join(setting.StaticRootPath, "views"),
		IndentJSON: true,
		Delims:     macaron.Delims{Left: "[[", Right: "]]"},
	}))
	m.Get("/api/frontend/settings/", hs.GetFrontendSettings)

	return m, hs
}

func TestHTTPServer_GetFrontendSettings_hideVersionAnonyomus(t *testing.T) {
	type buildInfo struct {
		Version string `json:"version"`
		Commit  string `json:"commit"`
		Env     string `json:"env"`
	}
	type settings struct {
		BuildInfo buildInfo `json:"buildInfo"`
	}

	cfg := setting.NewCfg()
	cfg.Env = "testing"
	cfg.BuildVersion = "7.8.9"
	cfg.BuildCommit = "01234567"
	m, hs := setupTestEnvironment(t, cfg)

	req := httptest.NewRequest(http.MethodGet, "/api/frontend/settings", nil)

	// TODO: Remove
	setting.BuildVersion = cfg.BuildVersion
	setting.BuildCommit = cfg.BuildCommit
	setting.Env = cfg.Env

	tests := []struct {
		desc        string
		hideVersion bool
		expected    settings
	}{
		{
			desc:        "Not hiding version",
			hideVersion: false,
			expected: settings{
				BuildInfo: buildInfo{
					Version: setting.BuildVersion,
					Commit:  setting.BuildCommit,
					Env:     setting.Env,
				},
			},
		},
		{
			desc:        "Hiding version",
			hideVersion: true,
			expected: settings{
				BuildInfo: buildInfo{
					Version: "",
					Commit:  "",
					Env:     setting.Env,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			hs.Cfg.AnonymousHideVersion = test.hideVersion
			expected := test.expected

			recorder := httptest.NewRecorder()
			m.ServeHTTP(recorder, req)
			got := settings{}
			err := json.Unmarshal(recorder.Body.Bytes(), &got)
			require.NoError(t, err)
			require.GreaterOrEqual(t, 400, recorder.Code, "status codes higher than 400 indicate a failure")

			assert.EqualValues(t, expected, got)
		})
	}
}
