package dao

import (
	"ampl/src/config"
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConn *gorm.DB

func CloseDbConn() {
	sql, _ := DbConn.DB()
	sql.Close()
}

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

func createEnum(db *gorm.DB) error {
	err := db.Exec(`
	DO $$ 
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_type') THEN
			CREATE TYPE status_type AS ENUM ('pending', 'in-progress', 'completed');
		END IF;
	END $$;
	`).Error
	if err != nil {
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
	sqlDb, err := db.DB()
	if err != nil {
		return db, err
	}

	sqlDb.SetMaxOpenConns(taskDb.MaxOpen)
	sqlDb.SetMaxIdleConns(taskDb.MaxIdle)
	sqlDb.SetConnMaxLifetime(time.Duration(1) * time.Hour)

	createEnum(db)

	db.AutoMigrate(&Tasks{})
	return db, nil
}
