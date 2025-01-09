package orm

import (
	"ampl/src/cache"
	"fmt"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createDatabase() error {
	var taskDb = cache.Config.Db
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

func AutoMigration() error {
	var taskDb = cache.Config.Db
	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%d sslmode=disable TimeZone=Asia/Kolkata",
		taskDb.Host, taskDb.User, taskDb.Password, strings.ToLower(taskDb.Database), taskDb.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&Tasks{})
	return nil
}

func Migrate() error {
	var err error
	err = createDatabase()
	if err != nil {
		return err
	}
	err = AutoMigration()
	if err != nil {
		return err
	}
	return err
}
