package main

import (
	"Medqueue-Alta-BE/config"
	sd "Medqueue-Alta-BE/features/schedule/data"
	sh "Medqueue-Alta-BE/features/schedule/handler"
	ss "Medqueue-Alta-BE/features/schedule/services"
	"Medqueue-Alta-BE/features/user/data"
	"Medqueue-Alta-BE/features/user/handler"
	"Medqueue-Alta-BE/features/user/services"
	"Medqueue-Alta-BE/helper"
	"Medqueue-Alta-BE/middlewares"
	"Medqueue-Alta-BE/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()            // inisiasi echo
	cfg := config.InitConfig() // baca seluruh system variable
	db := config.InitSQL(cfg)  // konek DB

	userData := data.New(db)
	userService := services.NewService(userData, helper.NewPasswordManager(), middlewares.NewMidlewareJWT())
	userHandler := handler.NewUserHandler(userService)

	scheduleData := sd.New(db)
	scheduleService := ss.NewScheduleService(scheduleData)
	scheduleHandler := sh.NewHandler(scheduleService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	routes.InitRoute(e, userHandler, scheduleHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
