package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name     string `gorm:"notnull"`
	Email    string `gorm:"unique;notnull"`
	Password string `gorm:"notnull"`
}
