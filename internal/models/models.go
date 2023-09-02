package models

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Title     string
	Completed bool
}
