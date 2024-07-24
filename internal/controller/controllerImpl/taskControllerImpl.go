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

// CreateTask godoc
//
//	@Summary Create a new task
//	@Description Create a new task with the input payload
//	@Tags tasks
//	@Accept json
//	@Produce json
//	@Param task body dto.TaskRequestDto true "Task"
//	@Success 200 {object} dto.TaskResponseDto
//	@Failure 400 {object} error
//	@Router /todo-list/tasks [post]
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

// UpdateTask godoc
// @Summary Update an existing task
// @Description Update the task with the input payload
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body dto.TaskRequestDto true "Task"
// @Success 200 {string} string "ok"
// @Failure 400 {object} error
// @Router /todo-list/tasks/{id} [put]
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

// DeleteTask godoc
// @Summary Delete an existing task
// @Description Delete the task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {string} string "ok"
// @Failure 400 {object} error
// @Router /todo-list/tasks/{id} [delete]
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

// CompleteTaskById godoc
// @Summary Complete an existing task
// @Description Complete the task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {string} string "ok"
// @Failure 400 {object} error
// @Router /todo-list/tasks/{id}/complete [put]
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

// GetTasksByStatus godoc
// @Summary Get tasks by status
// @Description Get tasks filtered by their status
// @Tags tasks
// @Accept json
// @Produce json
// @Param status path string true "Task status"
// @Success 200 {array} dto.TaskResponseDto
// @Router /todo-list/tasks [get]
func (t *taskController) GetTasksByStatus(ctx *gin.Context) {
	status := ctx.Query("status")

	tasks := t.Service.GetTasksByStatus(status)

	ctx.JSON(http.StatusOK, tasks)

}
