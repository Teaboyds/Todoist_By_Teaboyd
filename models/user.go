package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `json:"username" gorm:"unique;not null"`
	Password string    `json:"password" gorm:"not null"`
	Email    string    `json:"email" gorm:"unique;not null"`
	Projects []Project `gorm:"foreignKey:UserID" json:"-"`
	Tasks    []Task    `gorm:"foreignKey:UserID" json:"-"`
}
