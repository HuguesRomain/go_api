package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email      string `gorm:"not null;unique_index"`
	Password   string `gorm:"not null"`
	IsMerchant bool   `gorm:"not null;default:false"`
	Shops      []Shop
}