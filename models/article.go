package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model 
	Id int `json:"id,omitempty,autoIncrement"` 
	Content string `json:"content" gorm:"type:longtext" validate:"required"`
	IsDone int `json:"is_done" validate:"required"` 
	CourseId int `json:"course_id" validate:"required"` 
}