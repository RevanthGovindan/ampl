package tests

import (
	"ampl/src/config"
	"ampl/src/controllers"
	"ampl/src/dao"
	"ampl/src/models"
	"ampl/src/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupConfigs(t *testing.T) *gin.Engine {
	err := utils.InitializeConfigs(&config.Config)
	assert.NoError(t, err)
	err = dao.RedisConn.Init(config.Config.Redis)
	assert.NoError(t, err)

	dao.DbConn, err = dao.InitializeDb()
	assert.NoError(t, err)

	config.CloudPrivateKey, err = utils.LoadPrivateKey(config.Config.PvtKeyPath)
	assert.NoError(t, err)

	config.CloudPublicKey, err = utils.LoadPublicKey(config.Config.PubKeyPath)
	assert.NoError(t, err)

	utils.InitLogging(config.Config.Log)

	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	var router = controllers.Router{}
	router.RegisterRoutes(engine)
	return engine
}

func makeRequest(router *gin.Engine, method, url string, body interface{}, headers map[string]string) (*http.Response, []byte, error) {
	var reqBody []byte
	if body != nil {
		reqBody, _ = json.Marshal(body)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	return recorder.Result(), recorder.Body.Bytes(), nil
}

func reusableLogin(engine *gin.Engine, t *testing.T) models.LoginResponse {
	taskRequest := models.LoginRequest{
		Name:     "ampl",
		Password: "amplampl",
	}

	reqBody, err := json.Marshal(taskRequest)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "/public/login", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var loginResp models.LoginResponse
	err = json.Unmarshal(w.Body.Bytes(), &loginResp)
	if err != nil {
		t.Fatal("Failed to parse response body:", err)
	}
	return loginResp
}

func TestLogin(t *testing.T) {
	engine := setupConfigs(t)
	loginResp := reusableLogin(engine, t)
	assert.Contains(t, loginResp.Name, "ampl")
	assert.Contains(t, loginResp.Type, "Bearer")
}

func TestCreateTask(t *testing.T) {
	engine := setupConfigs(t)
	loginResp := reusableLogin(engine, t)
	assert.NotEmpty(t, loginResp.Token)
	taskRequest := models.CreateTask{
		Title:       "Test Task",
		Description: "Test Task Description",
	}

	resp, resBytes, err := makeRequest(engine, http.MethodPost, "/tasks", taskRequest,
		map[string]string{"Authorization": fmt.Sprintf("Bearer %s", loginResp.Token)})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var taskResp dao.Tasks
	err = json.Unmarshal(resBytes, &taskResp)
	if err != nil {
		t.Fatal("Failed to parse response body:", err)
	}

	assert.Contains(t, taskResp.Title, "Test Task")
	assert.Contains(t, taskResp.Description, "Test Task Description")
}

func TestGetTaskById(t *testing.T) {
	engine := setupConfigs(t)
	loginResp := reusableLogin(engine, t)
	assert.NotEmpty(t, loginResp.Token)
	taskRequest := models.CreateTask{
		Title:       "Test Task",
		Description: "Test Task Description",
	}

	resp, resBytes, err := makeRequest(engine, http.MethodPost, "/tasks", taskRequest,
		map[string]string{"Authorization": fmt.Sprintf("Bearer %s", loginResp.Token)})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var taskResp dao.Tasks
	err = json.Unmarshal(resBytes, &taskResp)
	if err != nil {
		t.Fatal("Failed to parse response body:", err)
	}

	resp, resBytes, err = makeRequest(engine, http.MethodGet, fmt.Sprintf("/tasks/%d", taskResp.ID), nil,
		map[string]string{"Authorization": fmt.Sprintf("Bearer %s", loginResp.Token)})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var getTaskResp dao.Tasks
	err = json.Unmarshal(resBytes, &getTaskResp)
	if err != nil {
		t.Fatal("Failed to parse response body:", err)
	}

	assert.Contains(t, getTaskResp.Title, "Test Task")
	assert.Contains(t, getTaskResp.Description, "Test Task Description")
}

func TestDeleteTaskById(t *testing.T) {
	engine := setupConfigs(t)
	loginResp := reusableLogin(engine, t)
	assert.NotEmpty(t, loginResp.Token)
	taskRequest := models.CreateTask{
		Title:       "Test Task",
		Description: "Test Task Description",
	}

	resp, resBytes, err := makeRequest(engine, http.MethodPost, "/tasks", taskRequest,
		map[string]string{"Authorization": fmt.Sprintf("Bearer %s", loginResp.Token)})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var taskResp dao.Tasks
	err = json.Unmarshal(resBytes, &taskResp)
	if err != nil {
		t.Fatal("Failed to parse response body:", err)
	}

	resp, _, err = makeRequest(engine, http.MethodDelete, fmt.Sprintf("/tasks/%d", taskResp.ID), nil,
		map[string]string{"Authorization": fmt.Sprintf("Bearer %s", loginResp.Token)})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	resp, _, err = makeRequest(engine, http.MethodGet, fmt.Sprintf("/tasks/%d", taskResp.ID), nil,
		map[string]string{"Authorization": fmt.Sprintf("Bearer %s", loginResp.Token)})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestUpdateTaskById(t *testing.T) {
	engine := setupConfigs(t)
	loginResp := reusableLogin(engine, t)
	assert.NotEmpty(t, loginResp.Token)
	taskRequest := models.CreateTask{
		Title:       "Test Task",
		Description: "Test Task Description",
	}

	resp, resBytes, err := makeRequest(engine, http.MethodPost, "/tasks", taskRequest,
		map[string]string{"Authorization": fmt.Sprintf("Bearer %s", loginResp.Token)})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var taskResp dao.Tasks
	err = json.Unmarshal(resBytes, &taskResp)
	if err != nil {
		t.Fatal("Failed to parse response body:", err)
	}

	updateRequest := models.UpdateTask{
		CreateTask: models.CreateTask{
			Title:       "Test update",
			Description: "Test update",
		},
		Status: "completed",
	}

	resp, resBytes, err = makeRequest(engine, http.MethodPut, fmt.Sprintf("/tasks/%d", taskResp.ID), updateRequest,
		map[string]string{"Authorization": fmt.Sprintf("Bearer %s", loginResp.Token)})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var upTaskResp dao.Tasks
	err = json.Unmarshal(resBytes, &upTaskResp)
	if err != nil {
		t.Fatal("Failed to parse response body:", err)
	}

	assert.Contains(t, upTaskResp.Title, "Test update")
	assert.Contains(t, upTaskResp.Description, "Test update")
}
