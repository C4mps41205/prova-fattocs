package repository

import (
	"database/sql"
	"errors"
	"log/slog"
	"prova-fattocs/internal/domain"
)

type TaskRepository interface {
	List() ([]domain.Task, error)
	Create(task domain.Task) error
	Update(id int64, task domain.Task) error
	Delete(id int64) error
	Reorder(id int64, direction int64) error
	ExistsByName(name string, id int64) (bool, error)
}

type PostgresTaskRepository struct {
	db *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) TaskRepository {
	slog.Info("Creating new PostgresTaskRepository")
	return &PostgresTaskRepository{db: db}
}

func (r *PostgresTaskRepository) List() ([]domain.Task, error) {
	slog.Info("Listing all tasks")
	rows, err := r.db.Query("SELECT id, name, cost, deadline, presentation_order FROM tasks ORDER BY presentation_order")
	if err != nil {
		slog.Error("Failed to query tasks", "error", err)
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			slog.Error("Failed to close rows", "error", err)
		}
	}()

	var tasks []domain.Task
	for rows.Next() {
		var t domain.Task
		err := rows.Scan(&t.ID, &t.Name, &t.Cost, &t.Deadline, &t.OrderNumber)
		if err != nil {
			slog.Error("Failed to scan task row", "error", err)
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		slog.Error("Error iterating rows", "error", err)
		return nil, err
	}

	slog.Info("Successfully listed tasks", "count", len(tasks))
	return tasks, nil
}

func (r *PostgresTaskRepository) Create(task domain.Task) error {
	slog.Info("Creating new task", "name", task.Name, "cost", task.Cost, "deadline", task.Deadline)

	exists, err := r.ExistsByName(task.Name, task.ID)
	if err != nil {
		slog.Error("Failed to check if task exists", "name", task.Name, "error", err)
		return err
	}
	if exists {
		slog.Warn("Task with this name already exists", "name", task.Name)
		return errors.New("task with this name already exists")
	}

	var maxOrder int
	err = r.db.QueryRow("SELECT COALESCE(MAX(presentation_order), 0) FROM tasks").Scan(&maxOrder)
	if err != nil {
		slog.Error("Failed to get max presentation order", "error", err)
		return err
	}
	task.OrderNumber = maxOrder + 1

	result, err := r.db.Exec("INSERT INTO tasks (name, cost, deadline, presentation_order) VALUES ($1, $2, $3, $4)",
		task.Name, task.Cost, task.Deadline, task.OrderNumber)
	if err != nil {
		slog.Error("Failed to insert task", "name", task.Name, "error", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("Failed to get rows affected", "error", err)
	} else {
		slog.Info("Task created successfully", "name", task.Name, "rows_affected", rowsAffected)
	}

	return err
}

func (r *PostgresTaskRepository) Update(id int64, task domain.Task) error {
	slog.Info("Updating task", "id", id, "name", task.Name, "cost", task.Cost, "deadline", task.Deadline)

	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM tasks WHERE name=$1 AND id<>$2)", task.Name, id).Scan(&exists)
	if err != nil {
		slog.Error("Failed to check if another task with same name exists", "name", task.Name, "id", id, "error", err)
		return err
	}
	if exists {
		slog.Warn("Task with this name already exists", "name", task.Name)
		return errors.New("task with this name already exists")
	}

	result, err := r.db.Exec("UPDATE tasks SET name=$1, cost=$2, deadline=$3 WHERE id=$4",
		task.Name, task.Cost, task.Deadline, id)
	if err != nil {
		slog.Error("Failed to update task", "id", id, "error", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("Failed to get rows affected", "error", err)
	} else {
		slog.Info("Task updated successfully", "id", id, "rows_affected", rowsAffected)
	}

	return err
}

func (r *PostgresTaskRepository) Delete(id int64) error {
	slog.Info("Deleting task", "id", id)

	result, err := r.db.Exec("DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		slog.Error("Failed to delete task", "id", id, "error", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("Failed to get rows affected", "error", err)
	} else {
		slog.Info("Task deleted successfully", "id", id, "rows_affected", rowsAffected)
	}

	return err
}

func (r *PostgresTaskRepository) ExistsByName(name string, id int64) (bool, error) {
	slog.Debug("Checking if task exists by name", "name", name, "id", id)

	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM tasks WHERE name=$1 AND id<>$2)", name, id).Scan(&exists)
	if err != nil {
		slog.Error("Failed to check if task exists", "name", name, "error", err)
		return false, err
	}

	slog.Debug("Task existence check completed", "name", name, "exists", exists)
	return exists, err
}

func (r *PostgresTaskRepository) Reorder(id, direction int64) error {
	slog.Info("Reordering task", "id", id, "direction", direction)

	var currentOrder int
	err := r.db.QueryRow("SELECT presentation_order FROM tasks WHERE id=$1", id).Scan(&currentOrder)
	if err != nil {
		slog.Error("Failed to get current presentation order", "id", id, "error", err)
		return err
	}

	swapOrder := direction

	var swapID int64
	err = r.db.QueryRow("SELECT id FROM tasks WHERE presentation_order=$1", swapOrder).Scan(&swapID)
	if err == sql.ErrNoRows {
		slog.Info("No task found to swap with", "swap_order", swapOrder)
		return nil
	} else if err != nil {
		slog.Error("Failed to get task for swapping", "swap_order", swapOrder, "error", err)
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		slog.Error("Failed to begin transaction", "error", err)
		return err
	}

	defer func() {
		if err != nil {
			slog.Info("Rolling back transaction")
			tx.Rollback()
		}
	}()

	_, err = tx.Exec("UPDATE tasks SET presentation_order=$1 WHERE id=$2", swapOrder, id)
	if err != nil {
		slog.Error("Failed to update presentation order for first task", "id", id, "order", swapOrder, "error", err)
		return err
	}

	_, err = tx.Exec("UPDATE tasks SET presentation_order=$1 WHERE id=$2", currentOrder, swapID)
	if err != nil {
		slog.Error("Failed to update presentation order for second task", "id", swapID, "order", currentOrder, "error", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		slog.Error("Failed to commit transaction", "error", err)
		return err
	}

	slog.Info("Tasks reordered successfully", "task1_id", id, "task2_id", swapID)
	return nil
}
