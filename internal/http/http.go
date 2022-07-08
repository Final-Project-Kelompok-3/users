package http

import (
	"github.com/Final-Project-Kelompok-3/users/internal/app/role"
	"github.com/Final-Project-Kelompok-3/users/internal/app/user"
	"github.com/Final-Project-Kelompok-3/users/internal/factory"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {

	v1 := e.Group("user/v1")
	role.NewHandler(f).Route(v1.Group("/roles"))
	user.NewHandler(f).Route(v1.Group("/users"))
}
