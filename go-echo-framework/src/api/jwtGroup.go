package api

import (
	"api/handlers"

	"github.com/labstack/echo"
)

func JwtGroup(g *echo.Group) {

	// localhost:3000/jwt/main
	g.GET("/main", handlers.MainJwt)
}
