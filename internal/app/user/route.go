package user

import (
	"github.com/Final-Project-Kelompok-3/users/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("/register", h.Register, middleware.AuthApiKey)
	g.POST("/login", h.Login, middleware.AuthApiKey)

	g.GET("", h.Get, middleware.AuthApiKey)
	g.GET("/:id", h.GetByID, middleware.AuthApiKey)
	g.POST("", h.Create, middleware.AuthJWT)
	g.PUT("/:id", h.Update, middleware.AuthJWT)
	g.DELETE("/:id", h.Delete, middleware.AuthJWT)
	
}
