package controllers

import (
	"ampl/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct{}

func (Router) SetupRoutes() *gin.Engine {
	r := gin.Default()

	if !utils.IsRelease() {
		r.StaticFS("/docs", http.Dir("./docs"))
	}
	public := r.Group("/public")
	{
		public.GET("/tasks", getAllTasks)
	}

	authorized := r.Group("/")

	{
		authorized.GET("/tasks/{id}", getTaskById)
		authorized.POST("/tasks", createTask)
		authorized.PUT("/tasks/{id}", updateTaskById)
		authorized.DELETE("/tasks/{id}", deleteTaskById)
	}

	return r
}
