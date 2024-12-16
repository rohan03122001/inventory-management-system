package database

import (
	"fmt"

	"github.com/rohan03122001/inventory-management-system/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB *gorm.DB
}


func NewDatabase(config *config.Config) (*Database, error){
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        config.DB_HOST,
        config.DB_USER,
        config.DB_PASSWORD,
        config.DB_NAME,
        config.DB_PORT,
    )

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info), // Set logging level
    })

	if err!= nil{
		fmt.Errorf("Error Connecting To database %w",err)
	}

	//Test The connection

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Errorf("Error Getting database instance %w",err)
	}

	if err := sqlDB.Ping(); err !=nil{
		fmt.Errorf("Failed to ping database instance %w",err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return &Database{DB: db}, nil
}