package controllers

import (
	"ampl/src/config"
	"ampl/src/dao"
	"ampl/src/models"
	"ampl/src/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Router struct{}

func (f *Router) Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		var splits = strings.Split(authHeader, " ")
		if authHeader == "" || len(splits) < 2 || !strings.EqualFold(strings.TrimSpace(splits[0]), "bearer") {
			c.JSON(http.StatusUnauthorized, models.ErrResponse{Error: "unauthorized"})
			c.Abort()
			return
		}
		var token = strings.TrimSpace(splits[1])
		data, err := dao.RedisConn.GetTokenData(token)
		if err != nil || data == "" {
			c.JSON(http.StatusUnauthorized, models.ErrResponse{Error: "unauthorized"})
			c.Abort()
			return
		}
		claims, err := utils.JwtDecrypt(token, config.CloudPublicKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrResponse{Error: "unauthorized"})
			c.Abort()
			return
		}
		exp, exists := claims["exp"]
		if !exists || time.Now().After(time.UnixMilli(int64(exp.(float64)))) {
			c.JSON(http.StatusUnauthorized, models.ErrResponse{Error: "unauthorized"})
			c.Abort()
			return
		}
		c.Request.Header.Set("x-user-data", data)
		c.Next()
	}
}

func (f *Router) SetupRoutes() *gin.Engine {
	r := gin.Default()

	if !utils.IsRelease() {
		r.StaticFS("/docs", http.Dir("./docs"))
	}
	public := r.Group("/public")
	{
		public.GET("/tasks", getAllTasks)
		public.POST("/login", login)
	}

	authorized := r.Group("/")
	{
		authorized.Use(f.Authorized())
		authorized.GET("/tasks/:id", getTaskById)
		authorized.POST("/tasks", createTask)
		authorized.PUT("/tasks/:id", updateTaskById)
		authorized.DELETE("/tasks/:id", deleteTaskById)
	}

	return r
}
