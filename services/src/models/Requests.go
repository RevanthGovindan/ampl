package models

type LoginRequest struct {
	// required: true
	// example: ampl
	Name string `json:"name" binding:"required"`
	// required: true
	// example: amplampl
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type CreateTask struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateTask struct {
	CreateTask
	Status string `gorm:"type:status_type" json:"status" validate:"oneof=pending in-progress completed"`
}

type GetTaskParams struct {
	Limit int `form:"limit" validate:"gte=1"`
	Page  int `form:"pageNo" validate:"gte=1"`
}
