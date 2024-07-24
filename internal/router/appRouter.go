package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luciorim/todo-server/internal/controller"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/luciorim/todo-server/docs"
)

type AppRouter struct {
	TaskController controller.TaskController

	Router *gin.Engine
}

func NewAppRouter(tc controller.TaskController) *AppRouter {
	return &AppRouter{
		TaskController: tc,
	}
}

func (ap *AppRouter) InitRoutes() {
	r := gin.Default()

	api := r.Group("/api/todo-list")
	{
		api.POST("/tasks", ap.TaskController.CreateTask)
		api.PUT("/tasks/:id", ap.TaskController.UpdateTask)
		api.DELETE("tasks/:id", ap.TaskController.DeleteTask)
		api.PUT("/tasks/:id/complete", ap.TaskController.CompleteTaskById)
		api.GET("/tasks", ap.TaskController.GetTasksByStatus)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ap.Router = r
}
