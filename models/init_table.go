package models

import "gorm.io/gorm"

func InitTableCustomer(db *gorm.DB) {
	db.Debug().AutoMigrate(&Customer{})
}
