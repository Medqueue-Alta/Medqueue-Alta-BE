package routes

import (
	"Medqueue-Alta-BE/config"
	reservation "Medqueue-Alta-BE/features/reservation"
	schedule "Medqueue-Alta-BE/features/schedule"
	user "Medqueue-Alta-BE/features/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ctl user.UserController, rc reservation.ReservationController, sc schedule.ScheduleController) {
	userRoute(c, ctl)
	reservationRoute(c, rc)
	scheduleRoute(c, sc)
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

func reservationRoute(c *echo.Echo, rc reservation.ReservationController) {
	c.POST("/reservations", rc.Add(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.GET("/reservations", rc.ShowMyReservation(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.PUT("/reservations/:id", rc.Update(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.DELETE("/reservations/:id", rc.Delete(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}

func scheduleRoute(c *echo.Echo, sc schedule.ScheduleController) {
	c.POST("/schedules", sc.Add(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))

	c.GET("/schedules", sc.ShowAllSchedules())

	c.GET("/schedules/poli", sc.ShowSchedulesByPoliID())

	c.GET("/schedules/:id", sc.ShowScheduleByID())

	c.PUT("/schedules/:id", sc.Update(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.DELETE("/schedules/:id", sc.Delete(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}