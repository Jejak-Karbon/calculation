package http

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/app/category_carbon_producer"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/app/carbon_producer"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	category_carbon_producer.NewHandler(f).Route(e.Group("/categories_carbon_producer"))
	carbon_producer.NewHandler(f).Route(e.Group("/carbon_producer"))
}
