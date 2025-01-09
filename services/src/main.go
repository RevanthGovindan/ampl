package main

import (
	"ampl/src/cache"
	"ampl/src/orm"
	"ampl/src/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	// var route = controllers.Router{}
	// var engine = route.SetupRoutes()
	utils.InitializeConfigs(&cache.Config)
	// engine.Run(":8000")
	err := orm.Migrate()
	fmt.Println(err)
}
