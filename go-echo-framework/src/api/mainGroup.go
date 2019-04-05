package api

import (
	"api/handlers"

	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	e.GET("/", handlers.Hello)
	e.GET("/login", handlers.Login)
	e.GET("/jwt/login", handlers.JwtLogin)
	e.GET("/books/:data", handlers.GetBooks)
	e.POST("/books", handlers.AddBook)
	e.POST("/authors", handlers.AddAuthor)
	e.POST("/contacts", handlers.AddContact)
}
