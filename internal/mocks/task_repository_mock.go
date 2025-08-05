package mocks

import "prova-fattocs/internal/domain"

type TaskRepositoryMock struct {
	ListFunc         func() ([]domain.Task, error)
	CreateFunc       func(task domain.Task) error
	UpdateFunc       func(id int64, task domain.Task) error
	DeleteFunc       func(id int64) error
	ReorderFunc      func(id int64, direction int64) error
	ExistsByNameFunc func(name string, id int64) (bool, error)
}

func (m *TaskRepositoryMock) List() ([]domain.Task, error) {
	return m.ListFunc()
}

func (m *TaskRepositoryMock) Create(task domain.Task) error {
	return m.CreateFunc(task)
}

func (m *TaskRepositoryMock) Update(id int64, task domain.Task) error {
	return m.UpdateFunc(id, task)
}

func (m *TaskRepositoryMock) Delete(id int64) error {
	return m.DeleteFunc(id)
}

func (m *TaskRepositoryMock) Reorder(id int64, direction int64) error {
	return m.ReorderFunc(id, direction)
}

func (m *TaskRepositoryMock) ExistsByName(name string, id int64) (bool, error) {
	return m.ExistsByNameFunc(name, id)
}
