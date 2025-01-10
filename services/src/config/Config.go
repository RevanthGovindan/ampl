package config

import (
	"ampl/src/models"

	"gorm.io/gorm"
)

var (
	Config models.Config = models.Config{}
	DbConn *gorm.DB
)
