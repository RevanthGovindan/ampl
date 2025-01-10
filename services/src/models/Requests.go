package models

type LoginRequest struct {
	// required: true
	// example: ampl
	Name string `json:"name"`
	// required: true
	// example: ampl
	Password string `json:"password"`
}

type CreateTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTask struct {
	CreateTask
	Status string `gorm:"type:status_type" json:"status"`
}
