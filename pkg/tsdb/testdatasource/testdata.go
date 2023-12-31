package testdatasource

import (
	"net/http"

	"github.com/famarker/grafarg-plugin-sdk-go/backend"
	"github.com/famarker/grafarg-plugin-sdk-go/backend/datasource"
	"github.com/famarker/grafarg-plugin-sdk-go/backend/resource/httpadapter"
	"github.com/famarker/grafarg/pkg/infra/log"
	"github.com/famarker/grafarg/pkg/plugins/backendplugin"
	"github.com/famarker/grafarg/pkg/plugins/backendplugin/coreplugin"
	"github.com/famarker/grafarg/pkg/registry"
)

func init() {
	registry.RegisterService(&testDataPlugin{})
}

type testDataPlugin struct {
	BackendPluginManager backendplugin.Manager `inject:""`
	logger               log.Logger
	scenarios            map[string]*Scenario
	queryMux             *datasource.QueryTypeMux
}

func (p *testDataPlugin) Init() error {
	p.logger = log.New("tsdb.testdata")
	p.scenarios = map[string]*Scenario{}
	p.queryMux = datasource.NewQueryTypeMux()
	p.registerScenarios()
	resourceMux := http.NewServeMux()
	p.registerRoutes(resourceMux)
	factory := coreplugin.New(backend.ServeOpts{
		QueryDataHandler:    p.queryMux,
		CallResourceHandler: httpadapter.New(resourceMux),
	})
	err := p.BackendPluginManager.Register("testdata", factory)
	if err != nil {
		p.logger.Error("Failed to register plugin", "error", err)
	}
	return nil
}
