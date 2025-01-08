package config

import (
	"fmt"
	"shaksham/models"
	"shaksham/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDbConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		utils.GetEnv("DB_HOST"),
		utils.GetEnv("DB_USERNAME"),
		utils.GetEnv("DB_PASSWORD"),
		utils.GetEnv("DB_NAME"),
		"5432",
		"disable",
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB %v", err)
	}

	err = db.AutoMigrate(
		models.Job{},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to set auto migrate %v", err)
	}

	Db = db

	return db, nil
}
