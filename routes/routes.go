package routes

import (
	"Medqueue-Alta-BE/config"
	user "Medqueue-Alta-BE/features/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ctl user.UserController) {
	userRoute(c, ctl)
}

func userRoute(c *echo.Echo, ctl user.UserController) {
	c.POST("/register", ctl.Add())
	c.POST("/login", ctl.Login())
	c.GET("/users", ctl.Profile(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.PUT("/users", ctl.Update(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.DELETE("/users", ctl.Delete(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}
