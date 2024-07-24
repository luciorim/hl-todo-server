package main

import (
	"github.com/luciorim/todo-server/internal/cache"
	"github.com/luciorim/todo-server/internal/config"
	"github.com/luciorim/todo-server/internal/controller/controllerImpl"
	"github.com/luciorim/todo-server/internal/router"
	"github.com/luciorim/todo-server/internal/service/serviceImpl"
	"log"
)

// @title Todo Server
// @version 2.0
// @description Server allows manage your daily tasks

// @host localhost:8181
// @BasePath /api/
func main() {
	//init config
	cfg := config.MustInit()

	//init cache
	appCache := cache.NewCache()

	//init services
	taskService := serviceImpl.NewTaskService(appCache)

	//init controllers
	taskController := controllerImpl.NewTaskController(taskService)

	//init app router
	appRouter := router.NewAppRouter(taskController)
	appRouter.InitRoutes()

	//run server
	if err := appRouter.Router.Run(":" + cfg.HTTPServer.Port); err != nil {
		log.Fatalf("Failed to run server: %s", err)
	}

}
