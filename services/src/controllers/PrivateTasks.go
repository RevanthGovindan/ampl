package controllers

import (
	"ampl/src/config"
	"ampl/src/dao"
	"ampl/src/models"
	"ampl/src/service"
	"ampl/src/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	var result dao.Tasks
	var taskService *service.TaskService = &service.TaskService{Db: dao.DbConn}
	var intId uint64
	intId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Error: err.Error()})
		return
	}
	result, err = taskService.GetTaskById(intId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Error: err.Error()})
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
	task.Status = utils.STATUS_PENDING
	var taskService service.TaskService = service.TaskService{Db: dao.DbConn}
	err = taskService.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, models.MsgResponse{Msg: "Created"})
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

	if err := config.Validate.Struct(req); err != nil {
		var errStr string
		for _, e := range err.(validator.ValidationErrors) {
			errStr += fmt.Sprintf("Validation failed for %s: %s", e.Field(), e.Tag())
		}
		c.JSON(http.StatusBadRequest, models.ErrResponse{Error: errStr})
		return
	}

	var taskService service.TaskService = service.TaskService{Db: dao.DbConn}
	var task dao.Tasks = dao.Tasks{Title: req.Title, Description: req.Description, ID: uint64(intId), Status: req.Status}
	_, err = taskService.UpdateTaskById(task)
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
	intId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Error: err.Error()})
		return
	}
	var taskService service.TaskService = service.TaskService{Db: dao.DbConn}
	err = taskService.DeleteTaskById(intId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.MsgResponse{Msg: "Deleted"})
}
