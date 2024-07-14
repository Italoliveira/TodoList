package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title     string `json:"title"`
	Details   string `json:"details"`
	Done      bool   `json:"done"`
	DeletedAt gorm.DeletedAt
}
