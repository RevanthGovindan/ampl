package controllers

import (
	"ampl/src/dao"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		Get All Task
// @Description	Get all the tasks
// @Tags        Tasks
// @Id 			get-task
// @Accept      json
// @Success		200  {object} []dao.Tasks
// @Produce     json
// @Router      /public/tasks [get]
func getAllTasks(c *gin.Context) {
	var results []dao.Tasks = make([]dao.Tasks, 0)
	err := dao.DbConn.GetAllTasks(&results)
	fmt.Println(err, results)
	c.JSON(http.StatusOK, results)
}
