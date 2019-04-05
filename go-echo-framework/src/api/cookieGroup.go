package api

import (
	"api/handlers"

	"github.com/labstack/echo"
)

func CookieGroup(g *echo.Group) {
	// localhost:3000/cookie/main
	g.GET("/main", handlers.MainCookie)
}
