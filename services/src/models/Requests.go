package models

type CreateTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTask struct {
	CreateTask
	ID     uint64 `gorm:"primaryKey" json:"id"`
	Status string `gorm:"type:status_type" json:"status"`
}
