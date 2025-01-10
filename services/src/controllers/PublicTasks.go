package controllers

import (
	"ampl/src/dao"
	"ampl/src/models"
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
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}
