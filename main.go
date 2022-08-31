package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/yumekiti/blog-api/config"
	"github.com/yumekiti/blog-api/infrastructure"
	"github.com/yumekiti/blog-api/interface/handler"
	"github.com/yumekiti/blog-api/usecase"
)

func main() {
	// repository
	taskRepository := infrastructure.NewTaskRepository(config.NewDB())
	// usecase
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	// handler
	taskHandler := handler.NewTaskHandler(taskUsecase)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	handler.InitRouting(e,
		taskHandler,
	)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
