package controllers

import (
	"ampl/src/dao"
	"ampl/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary		Get Task
// @Description	Get a specific task by id
// @Tags        Tasks
// @Id 			get-task-id
// @Accept      json
// @Success		200  {object} dao.Tasks
// @Produce     json
// @Param       id path string true "Id of the task"
// @Security 	http_bearer
// @Router      /tasks/{id} [get]
func getTaskById(c *gin.Context) {
	id := c.Param("id")
	var result dao.Tasks
	err := dao.DbConn.GetTaskById(id, &result)
	if err != nil {
		c.JSON(http.StatusOK, models.ErrResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary		Create
// @Description	Create a new task
// @Tags        Tasks
// @Id 			create-task
// @Accept      json
// @Success		200  {object} dao.Tasks
// @Produce     json
// @Param request body models.CreateTask true "Task data"
// @Security 	http_bearer
// @Router      /tasks [post]
func createTask(c *gin.Context) {
	var req models.CreateTask
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Error: err.Error()})
		return
	}
	var task dao.Tasks = dao.Tasks{Title: req.Title, Description: req.Description}
	task.Status = "pending"
	err = dao.DbConn.SaveTask(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Error: err.Error()})
		return
	}
}

// @Summary		Update Task
// @Description	update existing task
// @Tags        Tasks
// @Id 			update-task
// @Accept      json
// @Success		200  {object} dao.Tasks
// @Produce     json
// @Param request body models.UpdateTask true "Task data"
// @Param       id path string true "Id of the task"
// @Security 	http_bearer
// @Router      /tasks/{id} [put]
func updateTaskById(c *gin.Context) {
	var id = c.Param("id")
	intId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Error: err.Error()})
		return
	}
	var req models.UpdateTask
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Error: err.Error()})
		return
	}
	var task dao.Tasks = dao.Tasks{Title: req.Title, Description: req.Description, ID: uint64(intId), Status: req.Status}
	err = dao.DbConn.UpdateTaskById(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Error: err.Error()})
		return
	}
}

// @Summary		Delete task
// @Description	Delete task by id
// @Tags        Tasks
// @Id 			delete-task
// @Accept      json
// @Success		200  {object} models.MsgResponse
// @Produce     json
// @Param       id path string true "Id of the task"
// @Security 	http_bearer
// @Router      /tasks/{id} [delete]
func deleteTaskById(c *gin.Context) {
	var id = c.Param("id")
	err := dao.DbConn.DeleteTaskById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.MsgResponse{Msg: "Deleted"})
}
