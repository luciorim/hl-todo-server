package cache

import (
	"errors"
	"github.com/luciorim/todo-server/internal/entity"
	"github.com/patrickmn/go-cache"
)

type AppCache struct {
	Cache *cache.Cache
}

func NewCache() *AppCache {
	return &AppCache{
		Cache: cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (c *AppCache) Set(task *entity.Task) {
	c.Cache.Set(task.ID, task, cache.NoExpiration)
}

func (c *AppCache) Get(id string) (*entity.Task, error) {
	reqData, ok := c.Cache.Get(id)
	if !ok {
		return nil, errors.New("task with id: " + id + "not found")
	}

	return reqData.(*entity.Task), nil
}

func (c *AppCache) Delete(id string) error {
	_, ok := c.Cache.Get(id)
	if !ok {
		return errors.New("task with id: " + id + "not found")
	}

	c.Cache.Delete(id)
	return nil
}

func (c *AppCache) GetAllTasks() []*entity.Task {
	items := c.Cache.Items()
	tasks := make([]*entity.Task, 0, len(items))
	for _, item := range items {
		task, ok := item.Object.(*entity.Task)
		if ok {
			tasks = append(tasks, task)
		}
	}
	return tasks
}
