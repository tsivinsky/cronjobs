package db

import (
	"api/env"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func Init() error {
	var err error

	dsn := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", env.Env.DBHost, env.Env.DBUser, env.Env.DBPassword, env.Env.DBName)
	Db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		return err
	}

	err = Db.AutoMigrate()
	if err != nil {
		return err
	}

	return nil
}
