package routes

import (
	"database/sql"
	"prova-fattocs/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
	"prova-fattocs/internal/app"
	"prova-fattocs/internal/dto"
	"prova-fattocs/internal/infra/repository"
	"prova-fattocs/pkg/response"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	taskRepo := repository.NewPostgresTaskRepository(db)
	taskService := app.NewTaskService(taskRepo)

	// List all tasks
	// @Summary      Get all tasks
	// @Description  Returns all tasks
	// @Tags         Tasks
	// @Produce      json
	// @Success      200 {object} response.Response
	// @Failure      500 {object} response.Response
	// @Router       /tasks [get]
	r.GET("/tasks", func(c *gin.Context) {
		tasks, err := taskService.List()
		if err != nil {
			response.InternalServerError(c, "Failed to fetch tasks", nil)
			return
		}

		response.OK(c, "Tasks retrieved successfully", tasks)
	})

	// Create a task
	// @Summary      Create task
	// @Description  Creates a new task
	// @Tags         Tasks
	// @Accept       json
	// @Produce      json
	// @Param        task body dto.CreateTaskDTO true "Task payload"
	// @Success      201 {object} response.Response
	// @Failure      400 {object} response.Response
	// @Failure      500 {object} response.Response
	// @Router       /tasks [post]
	r.POST("/tasks", func(c *gin.Context) {
		var input dto.CreateTaskDTO
		if err := c.ShouldBindJSON(&input); err != nil {
			response.BadRequest(c, "Invalid input", nil)
			return
		}

		task := domain.Task{
			Name:     input.Name,
			Cost:     input.Cost,
			Deadline: input.Deadline,
		}

		err := taskService.Create(task)
		if err != nil {
			if err.Error() == "task with this name already exists" {
				response.BadRequest(c, err.Error(), nil)
				return
			}
			response.InternalServerError(c, "Failed to create task", nil)
			return
		}

		response.Created(c, "Task created successfully", task)
	})

	// Update a task
	// @Summary      Update task
	// @Description  Updates an existing task
	// @Tags         Tasks
	// @Accept       json
	// @Produce      json
	// @Param        id path int true "Task ID"
	// @Param        task body dto.UpdateTaskDTO true "Updated task data"
	// @Success      200 {object} response.Response
	// @Failure      400 {object} response.Response
	// @Failure      404 {object} response.Response
	// @Failure      500 {object} response.Response
	// @Router       /tasks/{id} [put]
	r.PUT("/tasks/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			response.BadRequest(c, "Invalid ID", nil)
			return
		}

		var input dto.UpdateTaskDTO
		if err := c.ShouldBindJSON(&input); err != nil {
			response.BadRequest(c, "Invalid input", nil)
			return
		}

		task := domain.Task{
			Name:     input.Name,
			Cost:     input.Cost,
			Deadline: input.Deadline,
		}

		err = taskService.Update(id, task)
		if err != nil {
			if err.Error() == "task with this name already exists" {
				response.BadRequest(c, err.Error(), nil)
				return
			}
			response.InternalServerError(c, "Failed to update task", nil)
			return
		}

		response.OK(c, "Task updated successfully", task)
	})

	// Delete a task
	// @Summary      Delete task
	// @Description  Deletes an existing task
	// @Tags         Tasks
	// @Produce      json
	// @Param        id path int true "Task ID"
	// @Success      200 {object} response.Response
	// @Failure      400 {object} response.Response
	// @Failure      500 {object} response.Response
	// @Router       /tasks/{id} [delete]
	r.DELETE("/tasks/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response.BadRequest(c, "Invalid ID", nil)
			return
		}

		if err := taskService.Delete(int64(id)); err != nil {
			response.InternalServerError(c, "Failed to delete task", nil)
			return
		}

		response.OK(c, "Task deleted successfully", nil)
	})

	// Reorder a task
	// @Summary      Reorder task
	// @Description  Changes the order of a task
	// @Tags         Tasks
	// @Accept       json
	// @Produce      json
	// @Param        id path int true "Task ID"
	// @Param        body body dto.ReorderTaskDTO true "New order"
	// @Success      200 {object} response.Response
	// @Failure      400 {object} response.Response
	// @Failure      500 {object} response.Response
	// @Router       /tasks/{id}/reorder [post]
	r.POST("/tasks/:id/reorder", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response.BadRequest(c, "Invalid ID", nil)
			return
		}

		var input dto.ReorderTaskDTO
		if err := c.ShouldBindJSON(&input); err != nil {
			response.BadRequest(c, "Invalid input", nil)
			return
		}

		if err := taskService.Reorder(int64(id), input.Order); err != nil {
			response.InternalServerError(c, "Failed to reorder task", nil)
			return
		}

		response.OK(c, "Task reordered successfully", nil)
	})
}
