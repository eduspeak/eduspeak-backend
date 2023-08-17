package models

import (
	"gorm.io/gorm"
	"time"
)

type Enroll struct {
	gorm.Model 
	Id int `json:"id,omitempty,autoIncrement"`  
	CourseId int `json:"course_id" validate:"required"` 
	Enrolled_at time.Time `json:"enrolled_at" validate:"required"` 
	UserId int `json:"user_id" validate:"required"` 
}