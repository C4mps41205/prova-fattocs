package app

import (
	"errors"
	"prova-fattocs/internal/domain"
	"prova-fattocs/internal/mocks"
	"testing"
)

func TestCreateTask_Success(t *testing.T) {
	mockRepo := &mocks.TaskRepositoryMock{
		ExistsByNameFunc: func(name string, id int64) (bool, error) {
			return false, nil
		},
		CreateFunc: func(task domain.Task) error {
			return nil
		},
	}

	service := NewTaskService(mockRepo)

	task := domain.Task{
		Name:     "New Task",
		Cost:     500,
		Deadline: "2025-08-10",
	}

	err := service.Create(task)
	if err != nil {
		t.Errorf("expected success but got error: %v", err)
	}
}

func TestCreateTask_DuplicateName(t *testing.T) {
	mockRepo := &mocks.TaskRepositoryMock{
		ExistsByNameFunc: func(name string, id int64) (bool, error) {
			return true, nil
		},
	}

	service := NewTaskService(mockRepo)

	task := domain.Task{
		Name:     "Duplicate",
		Cost:     100,
		Deadline: "2025-08-10",
	}

	err := service.Create(task)
	if err == nil {
		t.Error("expected duplicate name error but got nil")
	}
}

func TestUpdateTask_Success(t *testing.T) {
	mockRepo := &mocks.TaskRepositoryMock{
		ExistsByNameFunc: func(name string, id int64) (bool, error) {
			return false, nil
		},
		UpdateFunc: func(id int64, task domain.Task) error {
			return nil
		},
	}

	service := NewTaskService(mockRepo)

	task := domain.Task{
		Name:     "Updated Task",
		Cost:     250,
		Deadline: "2025-09-01",
	}

	err := service.Update(1, task)
	if err != nil {
		t.Errorf("expected success but got error: %v", err)
	}
}

func TestUpdateTask_DuplicateName(t *testing.T) {
	mockRepo := &mocks.TaskRepositoryMock{
		ExistsByNameFunc: func(name string, id int64) (bool, error) {
			return true, nil
		},
	}

	service := NewTaskService(mockRepo)

	task := domain.Task{
		Name:     "Duplicate Name",
		Cost:     150,
		Deadline: "2025-09-01",
	}

	err := service.Update(1, task)
	if err == nil {
		t.Error("expected duplicate name error but got nil")
	}
}

func TestGetAllTasks_Success(t *testing.T) {
	mockRepo := &mocks.TaskRepositoryMock{
		ListFunc: func() ([]domain.Task, error) {
			return []domain.Task{
				{ID: 1, Name: "Task 1", Cost: 100, Deadline: "2025-08-20", OrderNumber: 1},
				{ID: 2, Name: "Task 2", Cost: 200, Deadline: "2025-08-25", OrderNumber: 2},
			}, nil
		},
	}

	service := NewTaskService(mockRepo)

	tasks, err := service.List()
	if err != nil {
		t.Errorf("expected success but got error: %v", err)
	}
	if len(tasks) != 2 {
		t.Errorf("expected 2 tasks but got %d", len(tasks))
	}
}

func TestGetAllTasks_Error(t *testing.T) {
	mockRepo := &mocks.TaskRepositoryMock{
		ListFunc: func() ([]domain.Task, error) {
			return nil, errors.New("database error")
		},
	}

	service := NewTaskService(mockRepo)

	_, err := service.List()
	if err == nil {
		t.Error("expected error but got nil")
	}
}

func TestDeleteTask_Success(t *testing.T) {
	mockRepo := &mocks.TaskRepositoryMock{
		DeleteFunc: func(id int64) error {
			return nil
		},
	}

	service := NewTaskService(mockRepo)

	err := service.Delete(1)
	if err != nil {
		t.Errorf("expected success but got error: %v", err)
	}
}

func TestDeleteTask_Error(t *testing.T) {
	mockRepo := &mocks.TaskRepositoryMock{
		DeleteFunc: func(id int64) error {
			return errors.New("delete failed")
		},
	}

	service := NewTaskService(mockRepo)

	err := service.Delete(1)
	if err == nil {
		t.Error("expected error but got nil")
	}
}

func TestReorderTask_Success(t *testing.T) {
	mockRepo := &mocks.TaskRepositoryMock{
		ReorderFunc: func(id int64, direction int64) error {
			return nil
		},
	}

	service := NewTaskService(mockRepo)

	err := service.Reorder(1, 1)
	if err != nil {
		t.Errorf("expected success but got error: %v", err)
	}

	err = service.Reorder(1, -1)
	if err != nil {
		t.Errorf("expected success but got error: %v", err)
	}
}

func TestReorderTask_Error(t *testing.T) {
	mockRepo := &mocks.TaskRepositoryMock{
		ReorderFunc: func(id int64, direction int64) error {
			return errors.New("reorder failed")
		},
	}

	service := NewTaskService(mockRepo)

	err := service.Reorder(1, 1)
	if err == nil {
		t.Error("expected error but got nil")
	}
}
