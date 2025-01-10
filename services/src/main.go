package main

import (
	"ampl/src/config"
	"ampl/src/controllers"
	"ampl/src/orm"
	"ampl/src/utils"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// @title						Tasks Services
// @version						1.0
// @description					Module to manage tasks
// @tag.name					Tasks
// @tag.description				Tasks services
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
	config.DbConn, err = orm.InitializeDb()
	engine.Run(":8000")
	fmt.Println(err)
}
