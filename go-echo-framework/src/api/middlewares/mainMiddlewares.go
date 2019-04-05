package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetMainMiddlewares(e *echo.Echo) {

	e.Use(middleware.BodyLimit("2M"))

	// custom middlewares
	e.Use(serverHeader)

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "./static",
	}))

}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "GolangEcho/1.0")
		c.Response().Header().Set("FakeHeader", "ZzzZzzZ")
		return next(c)
	}
}
