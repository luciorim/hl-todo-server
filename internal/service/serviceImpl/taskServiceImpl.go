package serviceImpl

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/luciorim/todo-server/internal/cache"
	"github.com/luciorim/todo-server/internal/dto"
	"github.com/luciorim/todo-server/internal/entity"
	"github.com/luciorim/todo-server/internal/mapper"
	"github.com/luciorim/todo-server/internal/service"
	"time"
)

type taskService struct {
	Cache      *cache.AppCache
	TaskMapper mapper.TaskMapper
}

func NewTaskService(c *cache.AppCache) service.TaskService {
	return &taskService{
		Cache:      c,
		TaskMapper: mapper.NewTaskMapper(),
	}
}

func (t *taskService) CreateTask(req *dto.TaskRequestDto) (res *dto.TaskResponseDto, error error) {

	if err := req.ValidateRequest(); err != nil {
		return nil, err
	}

	const dateFormat = "2006-01-02"
	activeAt, err := time.Parse(dateFormat, req.ActiveAt)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	task := &entity.Task{
		ID:       uuid.NewString(),
		Title:    req.Title,
		ActiveAt: activeAt,
		Status:   false,
	}

	t.Cache.Set(task)

	return t.TaskMapper.MapToDto(task), nil
}

func (t *taskService) UpdateTask(id string, req *dto.TaskRequestDto) (error error) {

	if err := req.ValidateRequest(); err != nil {
		return err
	}

	taskEntity, err := t.Cache.Get(id)
	if err != nil {
		return errors.New(fmt.Sprintf("Task with id %s not found", id))
	}

	const dateFormat = "2006-01-02"
	activeAt, err := time.Parse(dateFormat, req.ActiveAt)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	taskEntity.Title = req.Title
	taskEntity.ActiveAt = activeAt

	t.Cache.Set(taskEntity)

	return nil
}

func (t *taskService) DeleteTask(id string) (error error) {

	err := t.Cache.Delete(id)
	if err != nil {
		return err
	}

	return nil

}

func (t *taskService) CompleteTaskById(id string) (error error) {
	taskEntity, err := t.Cache.Get(id)
	if err != nil {
		return errors.New(fmt.Sprintf("Task with id %s not found", id))
	}

	if taskEntity.Status == true {
		return errors.New(fmt.Sprintf("Task with id %s already done", id))
	}

	taskEntity.Status = true

	return nil
}

func (t *taskService) GetTasksByStatus(status string) []*dto.TaskResponseDto {
	done := false

	if status == "done" {
		done = true
	}

	tasks := t.Cache.GetAllTasks()

	var filteredTasks []*entity.Task

	for _, task := range tasks {
		if task.Status == done {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return t.TaskMapper.MapToDtos(filteredTasks)

}
