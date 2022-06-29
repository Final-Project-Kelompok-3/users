package http

import (
	"github.com/Final-Project-Kelompok-3/authentications/internal/app/role"
	"github.com/Final-Project-Kelompok-3/authentications/internal/app/user"
	"github.com/Final-Project-Kelompok-3/authentications/internal/factory"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {

	role.NewHandler(f).Route(e.Group("/roles"))
	user.NewHandler(f).Route(e.Group("/users"))
}
