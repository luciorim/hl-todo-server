package mapper

import (
	"github.com/luciorim/todo-server/internal/dto"
	"github.com/luciorim/todo-server/internal/entity"
)

type TaskMapper interface {
	MapToDto(ent *entity.Task) *dto.TaskResponseDto
	MapToDtos(ents []*entity.Task) []*dto.TaskResponseDto
}

type TaskMapperImpl struct{}

func NewTaskMapper() TaskMapper {
	return &TaskMapperImpl{}
}

func (t *TaskMapperImpl) MapToDto(ent *entity.Task) *dto.TaskResponseDto {
	res := &dto.TaskResponseDto{
		ID:       ent.ID,
		Title:    ent.Title,
		Done:     ent.Status,
		ActiveAt: ent.ActiveAt,
	}

	return res
}

func (t *TaskMapperImpl) MapToDtos(ents []*entity.Task) []*dto.TaskResponseDto {
	res := make([]*dto.TaskResponseDto, 0)
	for _, ent := range ents {
		res = append(res, t.MapToDto(ent))
	}

	return res
}
