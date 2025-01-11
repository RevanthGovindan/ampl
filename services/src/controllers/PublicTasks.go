package controllers

import (
	"ampl/src/config"
	"ampl/src/dao"
	"ampl/src/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary		Get All Task
// @Description	Get all the tasks
// @Tags        Tasks
// @Id 			get-task
// @Accept      json
// @Success		200  {object} []dao.Tasks
// @Produce     json
// @Param       pageNo query string false "Page no, if empty everything will sent in single shot"
// @Param       limit query string false "limit of records, if empty everything will sent in single shot"
// @Router      /public/tasks [get]
func getAllTasks(c *gin.Context) {
	var params models.GetTaskParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	if err := config.Validate.Struct(params); err != nil {
		var errStr string
		for _, e := range err.(validator.ValidationErrors) {
			errStr += fmt.Sprintf("Validation failed for %s: %s", e.Field(), e.Tag())
		}
		c.JSON(http.StatusBadRequest, models.ErrResponse{Error: errStr})
		return
	}
	var tasks []dao.Tasks = make([]dao.Tasks, 0)
	var total int64
	err := dao.DbConn.GetAllTasks(&tasks, params.Page, params.Limit, &total)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Error: err.Error()})
		return
	}
	var response models.AllTaskResponse[[]dao.Tasks] = models.AllTaskResponse[[]dao.Tasks]{Tasks: tasks, TotalCount: total}
	c.JSON(http.StatusOK, response)
}
