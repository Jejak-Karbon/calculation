package carbon_producer

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Get)
	g.POST("/:carbon_producer_id/calculate", h.Calculate,middleware.Authentication)
}
