package controllerImpl

import (
	"github.com/gin-gonic/gin"
	"github.com/luciorim/todo-server/internal/controller"
	"github.com/luciorim/todo-server/internal/dto"
	"github.com/luciorim/todo-server/internal/service"
	"net/http"
)

type taskController struct {
	Service service.TaskService
}

func NewTaskController(tc service.TaskService) controller.TaskController {
	return &taskController{
		Service: tc,
	}
}

func (t *taskController) CreateTask(ctx *gin.Context) {
	var taskReq *dto.TaskRequestDto

	if err := ctx.ShouldBindJSON(&taskReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskRes, err := t.Service.CreateTask(taskReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, taskRes)
}

func (t *taskController) UpdateTask(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is null"})
		return
	}

	var taskReq *dto.TaskRequestDto

	if err := ctx.ShouldBindJSON(&taskReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := t.Service.UpdateTask(id, taskReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "ok")

}

func (t *taskController) DeleteTask(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is null"})
		return
	}

	err := t.Service.DeleteTask(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "ok")

}

func (t *taskController) CompleteTaskById(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is null"})
		return
	}

	err := t.Service.CompleteTaskById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "ok")
}

func (t *taskController) GetTasksByStatus(ctx *gin.Context) {
	status := ctx.Param("status")

	tasks := t.Service.GetTasksByStatus(status)

	ctx.JSON(http.StatusOK, tasks)

}
