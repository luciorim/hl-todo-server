package main

import (
	"github.com/luciorim/todo-server/internal/cache"
	"github.com/luciorim/todo-server/internal/config"
	"github.com/luciorim/todo-server/internal/controller"
	"github.com/luciorim/todo-server/internal/controller/controllerImpl"
	"github.com/luciorim/todo-server/internal/service/serviceImpl"
	"log"
)

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
	appRouter := controller.NewAppRouter(taskController)

	appRouter.InitRoutes()

	//run server
	if err := appRouter.Router.Run(":" + cfg.HTTPServer.Port); err != nil {
		log.Fatalf("Failed to run server: %s", err)
	}

}
