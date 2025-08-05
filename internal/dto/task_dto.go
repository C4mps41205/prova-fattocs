package dto

type CreateTaskDTO struct {
	Name     string  `json:"name" binding:"required"`
	Cost     float64 `json:"cost" binding:"required"`
	Deadline string  `json:"deadline" binding:"required"`
}

type UpdateTaskDTO struct {
	Name     string  `json:"name" binding:"required"`
	Cost     float64 `json:"cost" binding:"required"`
	Deadline string  `json:"deadline" binding:"required"`
}

type ReorderTaskDTO struct {
	Order int64 `json:"order" binding:"required"`
}
