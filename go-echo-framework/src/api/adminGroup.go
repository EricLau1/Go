package api

import (
	"api/handlers"

	"github.com/labstack/echo"
)

func AdminGroup(g *echo.Group) {
	// localhost:3000/admin/main
	g.GET("/main", handlers.MainAdmin)
}
