package main

import (
	"Medqueue-BE/config"
	td "Medqueue-BE/features/todo/data"
	th "Medqueue-BE/features/todo/handler"
	ts "Medqueue-BE/features/todo/services"
	"Medqueue-BE/features/user/data"
	"Medqueue-BE/features/user/handler"
	"Medqueue-BE/features/user/services"
	"Medqueue-BE/routes"

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

	todoData := td.New(db)
	todoService := ts.NewTodoService(todoData)
	todoHandler := th.NewHandler(todoService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS()) // ini aja cukup
	routes.InitRoute(e, userHandler, todoHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
