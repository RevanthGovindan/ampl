package controllers

import (
	"ampl/src/orm"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		Get All Task
// @Description	Get all the tasks
// @Tags        Tasks
// @Id 			get-task
// @Accept      json
// @Success		200  {object} []orm.Tasks
// @Produce     json
// @Security 	http_bearer
// @Router      /public/tasks [get]
func getAllTasks(c *gin.Context) {
	var results []orm.Tasks = make([]orm.Tasks, 0)
	err := orm.GetAllTasks(&results)
	fmt.Println(err, results)
	c.JSON(http.StatusOK, results)
}
