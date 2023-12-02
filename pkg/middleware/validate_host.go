package middleware

import (
	"strings"

	"github.com/grafarg/grafarg/pkg/models"
	"github.com/grafarg/grafarg/pkg/setting"
	"gopkg.in/macaron.v1"
)

func ValidateHostHeader(cfg *setting.Cfg) macaron.Handler {
	return func(c *models.ReqContext) {
		// ignore local render calls
		if c.IsRenderCall {
			return
		}

		h := c.Req.Host
		if i := strings.Index(h, ":"); i >= 0 {
			h = h[:i]
		}

		if !strings.EqualFold(h, cfg.Domain) {
			c.Redirect(strings.TrimSuffix(cfg.AppURL, "/")+c.Req.RequestURI, 301)
			return
		}
	}
}