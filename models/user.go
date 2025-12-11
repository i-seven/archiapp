package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;size:255;not null"`
	Password string `grom:"size:255;not null"`
	Name     string `grom:"size:255"`
	Role     string `gorm:"default:'user'"`
}
