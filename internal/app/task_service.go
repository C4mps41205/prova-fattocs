package app

import (
	"errors"
	"prova-fattocs/internal/domain"
	"prova-fattocs/internal/infra/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) List() ([]domain.Task, error) {
	return s.repo.List()
}

func (s *TaskService) Create(task domain.Task) error {
	exists, err := s.repo.ExistsByName(task.Name, task.ID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("task with this name already exists")
	}
	return s.repo.Create(task)
}

func (s *TaskService) Update(id int64, updated domain.Task) error {
	exists, err := s.repo.ExistsByName(updated.Name, id)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("task with this name already exists")
	}
	return s.repo.Update(id, updated)
}

func (s *TaskService) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *TaskService) Reorder(id, direction int64) error {
	return s.repo.Reorder(id, direction)
}
