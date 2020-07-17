package main

import (
	"github.com/jinzhu/gorm"
)

// Migrations & Models and can be used as request validation schema

// Post Model
type Post struct {
	gorm.Model
	//column name				// validations
	Title        string  `gorm:"type:varchar(100);" json:"title" binding:"required"`
	Description        string  `gorm:"type:varchar(500);" json:"Description" binding:"required"`
	Status        string  `json:"status"`
}
