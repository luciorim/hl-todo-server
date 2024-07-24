package controller

import "github.com/gin-gonic/gin"

type AppRouter struct {
	TaskController TaskController

	Router *gin.Engine
}

func NewAppRouter(tc TaskController) *AppRouter {
	return &AppRouter{
		TaskController: tc,
	}
}

func (ap *AppRouter) InitRoutes() {
	r := gin.Default()

	api := r.Group("/api/todo-list")
	{
		api.POST("/tasks", ap.TaskController.CreateTask)
		api.PUT("/tasks/{id}", ap.TaskController.UpdateTask)
		api.DELETE("tasks/{id}", ap.TaskController.DeleteTask)
		api.PUT("/tasks", ap.TaskController.CompleteTaskById)
		api.GET("/tasks", ap.TaskController.GetTasksByStatus)
	}

	ap.Router = r
}
