package controller

import "github.com/gin-gonic/gin"

type TaskController interface {
	CreateTask(ctx *gin.Context)
	UpdateTask(ctx *gin.Context)
	DeleteTask(ctx *gin.Context)
	CompleteTaskById(ctx *gin.Context)
	GetTasksByStatus(ctx *gin.Context)
}
