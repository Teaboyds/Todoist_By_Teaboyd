package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Tasks  []Task `gorm:"foreignKey:ProjectID" json:"-"`
}
