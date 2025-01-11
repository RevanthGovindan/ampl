package controllers

import (
	"ampl/src/config"
	"ampl/src/dao"
	"ampl/src/models"
	"ampl/src/utils"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Router struct {
	limiters map[string]*rate.Limiter
	mu       sync.RWMutex
}

func (f *Router) GetLimiter(token string, rps int, burst int) *rate.Limiter {
	f.mu.RLock()
	limiter, exists := f.limiters[token]
	if exists {
		f.mu.RUnlock()
		return limiter
	}
	f.mu.RUnlock()
	f.mu.Lock()
	limiter = rate.NewLimiter(rate.Limit(rps), burst)
	f.limiters[token] = limiter
	f.mu.Unlock()
	return limiter
}

func (f *Router) requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Printf("Failed to read request body: %v", err)
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		latency := time.Since(startTime)
		log.Printf("%s %s %d %s %s\n",
			c.Request.URL.Path,
			c.Request.Method,
			c.Writer.Status(),
			latency, string(bodyBytes),
		)
		c.Next()
	}
}

func (f *Router) responseLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("%s %s %d\n",
			c.Request.RequestURI,
			c.Request.Method,
			c.Writer.Status(),
		)
		c.Next()
	}
}

func (f *Router) rateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		var splits = strings.Split(authHeader, " ")
		var token = strings.TrimSpace(splits[1])
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing access token"})
			c.Abort()
			return
		}

		limiter := f.GetLimiter(token, 5, 5)
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded. Try again later."})
			c.Abort()
			return
		}
		c.Next()
	}
}

func (f *Router) authorized() gin.HandlerFunc {
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

func (f *Router) RegisterRoutes(r *gin.Engine) {
	f.limiters = make(map[string]*rate.Limiter)
	r.Use(f.requestLogger())
	r.Use(f.responseLogger())
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
		authorized.Use(f.authorized())
		authorized.Use(f.rateLimiter())
		authorized.GET("/tasks/:id", getTaskById)
		authorized.POST("/tasks", createTask)
		authorized.PUT("/tasks/:id", updateTaskById)
		authorized.DELETE("/tasks/:id", deleteTaskById)
	}
}

func (f *Router) SetupRoutes() *gin.Engine {
	r := gin.Default()
	f.RegisterRoutes(r)
	return r
}
