package controllers

import (
	"ampl/src/models"
	"net/http"

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
	var response models.LoginResponse = models.LoginResponse{
		Name:  req.Name,
		Token: "111111111111111111111",
	}
	c.JSON(http.StatusBadRequest, response)
}
