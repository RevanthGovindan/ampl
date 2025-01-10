package controllers

import (
	"ampl/src/models"
	"ampl/src/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTaskById(c *gin.Context) {

}

// @Summary		Create
// @Description	Create a new task
// @Tags        Tasks
// @Id 			create-task
// @Accept      json
// @Success		200  {object} orm.Tasks
// @Produce     json
// @Param request body models.CreateTask true "Task data"
// @Security 	http_bearer
// @Router      /tasks [post]
func createTask(c *gin.Context) {
	var req models.CreateTask
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var task orm.Tasks = orm.Tasks{Title: req.Title, Description: req.Description}
	task.Status = "pending"
	err = orm.SaveTask(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
}

func updateTaskById(c *gin.Context) {

}

func deleteTaskById(c *gin.Context) {

}
