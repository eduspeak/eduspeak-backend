package models

import (
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model 
	Id int `json:"id,omitempty,autoIncrement"` 
	Title string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsDone int `json:"is_done" validate:"required"` 
	CourseId string `json:"course_id" validate:"required"` 
}