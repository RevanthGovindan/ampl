package models

type CreateTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTask struct {
	CreateTask
	Status string `gorm:"type:status_type" json:"status"`
}
