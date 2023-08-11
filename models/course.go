package models

import (
	"gorm.io/gorm"
	"time"
)

type Course struct {
	gorm.Model
	Id int `json:"id,omitempty,autoIncrement"` 
	Title string `json:"title" validate:"required"`
	GradeId int `json:"grade_id" validate:"required"`
	Description string `json:"description" validate:"required"`
	CoverCourse string `json:"cover_course" validate:"required"`
	IsPremium int `json:"is_premium" validate:"required"`
	Rating int `json:"rating" validate:"required"`
	IsDone int `json:"is_done" validate:"required"` 
	LastUpdate time.Time `json:"last_update" validate:"required"`
	TotalModule int `json:"total_module" validate:"required"`
}