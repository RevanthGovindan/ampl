package tests

import (
	"ampl/src/config"
	"ampl/src/dao"
	"ampl/src/service"
	"ampl/src/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnitCreateTask(t *testing.T) {
	err := utils.InitializeConfigs(&config.Config)
	assert.NoError(t, err)
	db, err := dao.InitializeDb()
	assert.NoError(t, err)
	taskService := &service.TaskService{Db: db}
	task := &dao.Tasks{Title: "Test Task", Description: "This is a test task", Status: utils.STATUS_PENDING}
	err = taskService.CreateTask(task)
	assert.NoError(t, err)
	var createdTask dao.Tasks
	db.First(&createdTask, task.ID)
	assert.Equal(t, task.Title, createdTask.Title)
	assert.Equal(t, task.Description, createdTask.Description)
	sql, err := db.DB()
	if err == nil {
		sql.Close()
	}
}

func TestGetTaskByID(t *testing.T) {
	err := utils.InitializeConfigs(&config.Config)
	assert.NoError(t, err)
	db, err := dao.InitializeDb()
	assert.NoError(t, err)
	taskService := &service.TaskService{Db: db}

	task := &dao.Tasks{Title: "Test Task", Description: "Test description", Status: utils.STATUS_PENDING}
	err = taskService.CreateTask(task)
	assert.NoError(t, err)

	fetchedTask, err := taskService.GetTaskById(task.ID)
	assert.NoError(t, err)
	assert.NotNil(t, fetchedTask)
	assert.Equal(t, task.ID, fetchedTask.ID)
	sql, err := db.DB()
	if err == nil {
		sql.Close()
	}
}

func TestUpdateTask(t *testing.T) {
	err := utils.InitializeConfigs(&config.Config)
	assert.NoError(t, err)
	db, err := dao.InitializeDb()
	assert.NoError(t, err)

	taskService := &service.TaskService{Db: db}

	task := &dao.Tasks{Title: "Test Task", Description: "Test description", Status: utils.STATUS_PENDING}
	err = taskService.CreateTask(task)
	assert.NoError(t, err)
	task.Title = "Updated Title"
	task.Description = "Updated description"
	task.Status = "completed"
	updatedTask, err := taskService.UpdateTaskById(*task)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Title", updatedTask.Title)
	assert.Equal(t, "Updated description", updatedTask.Description)
	assert.Equal(t, "completed", updatedTask.Status)
	sql, err := db.DB()
	if err == nil {
		sql.Close()
	}
}

func TestDeleteTask(t *testing.T) {
	db, err := dao.InitializeDb()
	assert.NoError(t, err)

	taskService := &service.TaskService{Db: db}
	task := &dao.Tasks{Title: "Test Task", Description: "Test description", Status: utils.STATUS_PENDING}
	err = taskService.CreateTask(task)
	assert.NoError(t, err)

	err = taskService.DeleteTaskById(task.ID)
	assert.NoError(t, err)

	var deletedTask dao.Tasks
	err = db.First(&deletedTask, task.ID).Error
	assert.Error(t, err)
	sql, err := db.DB()
	if err == nil {
		sql.Close()
	}
}
