package service

import "github.com/luciorim/todo-server/internal/dto"

type TaskService interface {
	CreateTask(req *dto.TaskRequestDto) (res *dto.TaskResponseDto, error error)
	UpdateTask(id string, req *dto.TaskRequestDto) (error error)
	DeleteTask(id string) (error error)
	CompleteTaskById(id string) (error error)
	GetTasksByStatus(status string) []*dto.TaskResponseDto
}
