package database

import (
	"fmt"
	"task_tracker/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.AppConfig) (*gorm.DB, error) {
	sdn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(sdn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v", err)
	}

	fmt.Println("Successfully connected!")

	return db, nil
}
