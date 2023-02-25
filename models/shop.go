package models

import "github.com/jinzhu/gorm"
type Shop struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Address  string `gorm:"not null"`
	UserID   uint
	IsActive bool `gorm:"not null;default:true"`
}
