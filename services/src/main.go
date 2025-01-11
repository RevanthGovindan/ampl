package main

import (
	"ampl/src/config"
	"ampl/src/controllers"
	"ampl/src/dao"
	"ampl/src/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func createConnections() {
	var err error
	dao.DbConn, err = dao.InitializeDb()
	if err != nil {
		panic(err)
	}
	err = dao.RedisConn.Init(config.Config.Redis)
	if err != nil {
		panic(err)
	}
}

func loadRsaCache() {
	var err error
	//load keys
	config.JwtRsaPrivateKey, err = utils.LoadPrivateKey(config.Config.PvtKeyPath)
	if err != nil {
		panic(err)
	}

	config.JwtRsaPublicKey, err = utils.LoadPublicKey(config.Config.PubKeyPath)
	if err != nil {
		panic(err)
	}
}

func shutdownhandler(quit chan os.Signal, server *http.Server) {
	<-quit
	log.Println("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//closing redis
	dao.RedisConn.Close()
	//closing postgres
	dao.CloseDbConn()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s", err)
	}
}

// @title						Tasks Services
// @version						1.0
// @description					Module to manage tasks
// @tag.name					Tasks
// @tag.description				Tasks services
// @tag.name					Login
// @tag.description				Login services
// @host 						localhost:8000
// @schemes 					http
// @BasePath					/
// @securityDefinitions.apikey	http_bearer
// @in 							header
// @name 						Authorization
func main() {
	var err error
	err = utils.InitializeConfigs(&config.Config)
	if err != nil {
		panic(err)
	}
	if utils.IsRelease() {
		gin.SetMode(gin.ReleaseMode)
	}
	//creating postgres and redis
	createConnections()
	loadRsaCache()

	var route = controllers.Router{}
	var engine = route.SetupRoutes()

	utils.InitLogging(config.Config.Log)

	server := &http.Server{
		Addr:    ":8000",
		Handler: engine,
	}

	go func() {
		fmt.Println("Starting server on :8000")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	shutdownhandler(quit, server)
	log.Println("Server exited")
}
