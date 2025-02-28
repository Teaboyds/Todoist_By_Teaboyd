package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	SetDate     time.Time `json:"set_date" gorm:"type:date;not null"`
	UserID      uint      `json:"user_id"`
	ProjectID   uint      `json:"project_id"`
}
