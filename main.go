package main

import (
	"Medqueue-Alta-BE/config"
	rd "Medqueue-Alta-BE/features/reservation/data"
	rh "Medqueue-Alta-BE/features/reservation/handler"
	rs "Medqueue-Alta-BE/features/reservation/services"
	"Medqueue-Alta-BE/features/user/data"
	"Medqueue-Alta-BE/features/user/handler"
	"Medqueue-Alta-BE/features/user/services"
	"Medqueue-Alta-BE/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()            // inisiasi echo
	cfg := config.InitConfig() // baca seluruh system variable
	db := config.InitSQL(cfg)  // konek DB

	userData := data.New(db)
	userService := services.NewService(userData)
	userHandler := handler.NewUserHandler(userService)

	reservationData := rd.New(db)
	reservationService := rs.NewTodoService(reservationData)
	reservationHandler := rh.NewHandler(reservationService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS()) 
	routes.InitRoute(e, userHandler, reservationHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
