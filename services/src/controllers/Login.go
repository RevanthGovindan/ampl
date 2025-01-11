package controllers

import (
	"ampl/src/config"
	"ampl/src/dao"
	"ampl/src/models"
	"ampl/src/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary		Login
// @Description	Login with name and password
// @Tags        Login
// @Id 			login
// @Accept      json
// @Success		200  {object} models.LoginResponse
// @Produce     json
// @Param request body models.LoginRequest true "Task data" example({"name":"ampl","password":"ampl"})
// @Router      /public/login [post]
func login(c *gin.Context) {
	var req models.LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Error: err.Error()})
		return
	}

	if strings.EqualFold(config.Config.Credentials.UserName, req.Name) &&
		strings.EqualFold(config.Config.Credentials.Password, req.Password) {
		var currTime = time.Now()
		var exp = time.Duration(1) * time.Hour
		token, _ := utils.JwtEncode(req.Name, currTime.Add(exp).UnixMilli(), config.CloudPrivateKey)
		var response models.LoginResponse = models.LoginResponse{
			Name: req.Name, Type: utils.TOKEN_TYPE, Token: token,
		}
		data, _ := json.Marshal(response)
		dao.RedisConn.SetToken(token, string(data), exp)
		c.JSON(http.StatusOK, response)
		return
	}
	fmt.Println("req")
	fmt.Println(req)
	c.JSON(http.StatusUnauthorized, models.ErrResponse{Error: "Invalid credentials"})
}
