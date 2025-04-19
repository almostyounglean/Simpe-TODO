package repository

import "simple-todo/internal/model"

type TaskRepository interface {
	Create(task *model.Task) error
	Update(*model.Task, error)
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetAll() ([]*model.Task, error)
}

type InMemoryTaskRepository struct {
	tasks  map[int]*model.Task
	nextID int
}

func CreateInMemoryRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks:  make(map[int]*model.Task),
		nextID: 1,
	}
}
