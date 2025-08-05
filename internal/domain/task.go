package domain

type Task struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Cost        float64 `json:"cost"`
	Deadline    string  `json:"deadline"`
	OrderNumber int     `json:"order_number"`
}
