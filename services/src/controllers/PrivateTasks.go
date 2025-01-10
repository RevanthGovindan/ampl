package controllers

import (
	"ampl/src/models"
	"ampl/src/orm"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		Get Task
// @Description	Get a specific task by id
// @Tags        Tasks
// @Id 			get-task-id
// @Accept      json
// @Success		200  {object} orm.Tasks
// @Produce     json
// @Param       id path string true "Id of the task"
// @Router      /tasks/{id} [get]
func getTaskById(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	var result orm.Tasks
	err := orm.GetTaskById(id, &result)
	if err != nil {
		c.JSON(http.StatusOK, models.ErrResponse{Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
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
