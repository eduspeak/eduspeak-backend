package models

import (
	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	Id       int `json:"id,omitempty,autoIncrement"`
	CourseId int `json:"course_id" validate:"required"`
	UserId   int `json:"user_id" validate:"required"`
	Rating   int `json:"rating" validate:"required" gorm:"default:0"`
	Feedback string `json:"feedback" gorm:"type:longtext" validate:"required"`
}
