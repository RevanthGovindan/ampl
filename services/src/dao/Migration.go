package dao

import (
	"ampl/src/config"
	"fmt"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDb() (*gorm.DB, error) {
	var err error
	err = createDatabase()
	if err != nil {
		return nil, err
	}
	db, err := autoMigration()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createDatabase() error {
	var taskDb = config.Config.Db
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d sslmode=disable TimeZone=Asia/Kolkata",
		taskDb.Host, taskDb.User, taskDb.Password, taskDb.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	defer func() {
		sql, _ := db.DB()
		sql.Close()
	}()

	var result int
	err = db.Raw(fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname = '%s'", strings.ToLower(taskDb.Database))).Scan(&result).Error
	if err != nil {
		return err
	}
	if result == 0 {
		query := fmt.Sprintf("CREATE DATABASE %s", strings.ToLower(taskDb.Database))
		err = db.Exec(query).Error
		return err
	}
	return nil
}

func autoMigration() (*gorm.DB, error) {
	var taskDb = config.Config.Db
	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%d sslmode=disable TimeZone=Asia/Kolkata",
		taskDb.Host, taskDb.User, taskDb.Password, strings.ToLower(taskDb.Database), taskDb.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Tasks{})
	return db, nil
}
