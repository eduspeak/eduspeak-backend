package models

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model 
	Id int `json:"id,omitempty,autoIncrement"` 
	VideoLink string `json:"video_link" validate:"required"`
	IsDone int `json:"is_done" validate:"required"` 
	CourseId int `json:"course_id" validate:"required"` 
}