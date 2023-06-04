package db

import (
	"fmt"

	"github.com/talesrc/taleco-store/customers/config"
	"github.com/talesrc/taleco-store/customers/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connect(connInfo config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s sslmode=disable",
		connInfo.Host, connInfo.Database, connInfo.User, connInfo.Password,
	)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func Setup() error {
	config := config.GetConfig()
	db, err := connect(config.Database)
	if err != nil {
		return fmt.Errorf("error when connecting to database %w", err)
	}
	if err = db.AutoMigrate(&models.Customer{}); err != nil {
		return fmt.Errorf("error when setup database %w", err)
	}
	return nil
}
