package main

import (
	"ampl/src/config"
	"ampl/src/controllers"
	"ampl/src/dao"
	"ampl/src/utils"
	"os"

	"github.com/gin-gonic/gin"
)

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
		os.Exit(1)
	}
	if utils.IsRelease() {
		gin.SetMode(gin.ReleaseMode)
	}
	var route = controllers.Router{}
	var engine = route.SetupRoutes()
	dao.DbConn.Db, err = dao.InitializeDb()
	if err != nil {
		os.Exit(1)
	}
	err = dao.RedisConn.Init(config.Config.Redis)
	if err != nil {
		os.Exit(1)
	}

	//load keys
	config.CloudPrivateKey, err = utils.LoadPrivateKey(config.Config.PvtKeyPath)
	if err != nil {
		os.Exit(1)
	}

	config.CloudPublicKey, err = utils.LoadPublicKey(config.Config.PubKeyPath)
	if err != nil {
		os.Exit(1)
	}
	utils.InitLogging(config.Config.Log)

	engine.Run(":8000")
}
