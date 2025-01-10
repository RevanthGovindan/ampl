package utils

import (
	"ampl/src/models"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func getLogFilePath(config models.LogInfo) string {
	logFilePath := config.File
	if err := os.MkdirAll(config.Dir, os.ModePerm); err == nil {
		logFilePath = filepath.Join(config.Dir, logFilePath)
	}
	return logFilePath
}

func InitLogging(config models.LogInfo) {
	file, err := os.OpenFile(getLogFilePath(config), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}
	log.SetOutput(file)
	gin.DefaultWriter = file
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}
